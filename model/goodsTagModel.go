package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ GoodsTagModel = (*customGoodsTagModel)(nil)
var GoodsTagTableName = "goods_tag"
var cacheGoodsTagUuidPrefix = "cache:goodsTag:uuid:"

type (
	// GoodsTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsTagModel.
	GoodsTagModel interface {
		goodsTagModel
		DeleteSoft(ctx context.Context, data *GoodsTag) error
		FindOneByUuid(ctx context.Context, uuid string) (*GoodsTag, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, goodsTag *GoodsTag) (resp []*GoodsTag, total int64, err error)
	}

	customGoodsTagModel struct {
		*defaultGoodsTagModel
	}
)

// NewGoodsTagModel returns a model for the database table.
func NewGoodsTagModel(conn sqlx.SqlConn, c cache.CacheConf) GoodsTagModel {
	//用于插入数据的字段
	goodsTagRowsExpectAutoSet = strings.Join(stringx.Remove(goodsTagFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	goodsTagRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsTagFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customGoodsTagModel{
		defaultGoodsTagModel: newGoodsTagModel(conn, c),
	}
}

func (m *customGoodsTagModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheGoodsTagUuidPrefix, uuid)
}

// 软删除
func (m *customGoodsTagModel) DeleteSoft(ctx context.Context, data *GoodsTag) error {
	goodsTagIdKey := m.formatPrimary(data.Id)
	goodsTagUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, goodsTagIdKey, goodsTagUuidKey)
	//完成
	return err
}

func (m *customGoodsTagModel) FindOneByUuid(ctx context.Context, uuid string) (*GoodsTag, error) {
	goodsTagUuidKey := m.formatUuidKey(uuid)
	var resp GoodsTag
	err := m.QueryRowCtx(ctx, &resp, goodsTagUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", goodsTagRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, uuid)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customGoodsTagModel) FindOne(ctx context.Context, id int64) (*GoodsTag, error) {
	goodsTagIdKey := m.formatPrimary(id)
	var resp GoodsTag
	err := m.QueryRowCtx(ctx, &resp, goodsTagIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ?  and `is_deleted` = '0' limit 1", goodsTagRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customGoodsTagModel) Insert(ctx context.Context, data *GoodsTag) (sql.Result, error) {
	goodsTagIdKey := m.formatPrimary(data.Id)
	goodsTagUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, goodsTagRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Status)
	}, goodsTagIdKey, goodsTagUuidKey)
	return ret, err
}

func (m *customGoodsTagModel) Update(ctx context.Context, data *GoodsTag) error {
	goodsTagIdKey := m.formatPrimary(data.Id)
	goodsTagUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, goodsTagRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Status, data.Id)
	}, goodsTagIdKey, goodsTagUuidKey)
	return err
}

//----------------------------------------------------------------

func (m *customGoodsTagModel) FindList(ctx context.Context, pageSize, page int64, keyword string, goodsTag *GoodsTag) (resp []*GoodsTag, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(goodsTagRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if goodsTag != nil {
		//name的优先级大于keyword
		if goodsTag.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", goodsTag.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", goodsTag.Name, "%"))
			mark = true
		}
		if goodsTag.Status != -1 {
			sq = sq.Where("`status` = ?", goodsTag.Status)
			sqCount = sqCount.Where("`status` = ?", goodsTag.Status)
			mark = true
		}
	}

	if !mark {
		if keyword = strings.TrimSpace(keyword); keyword != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", keyword, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", keyword, "%"))
		}
	}

	isPageQuery := pageSize > 0 && page > 0
	if isPageQuery {
		sq = sq.Offset(uint64((page - 1) * pageSize)).Limit(uint64(pageSize))
		queryCount, agrsCount, e := sqCount.ToSql()
		if e != nil {
			err = e
			return
		}
		if err = m.QueryRowNoCacheCtx(ctx, &total, queryCount, agrsCount...); err != nil {
			return
		}
	}

	query, agrs, err := sq.ToSql()
	if err != nil {
		return
	}
	if err = m.QueryRowsNoCacheCtx(ctx, &resp, query, agrs...); err != nil {
		return
	}
	return
}
