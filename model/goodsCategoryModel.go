/*
 * @Author: licat
 * @Date: 2023-01-10 15:16:02
 * @LastEditors: licat
 * @LastEditTime: 2023-01-10 15:56:32
 * @Description: licat233@gmail.com
 */
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

var _ GoodsCategoryModel = (*customGoodsCategoryModel)(nil)
var GoodsCategoryTableName = "goods_category"
var cacheGoodsCategoryUuidPrefix = "cache:goodsCategory:uuid:"

type (
	// GoodsCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsCategoryModel.
	GoodsCategoryModel interface {
		goodsCategoryModel
		DeleteSoft(ctx context.Context, data *GoodsCategory) error
		FindOneByUuid(ctx context.Context, uuid string) (*GoodsCategory, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, goodsCategory *GoodsCategory) (resp []*GoodsCategory, total int64, err error)
	}

	customGoodsCategoryModel struct {
		*defaultGoodsCategoryModel
	}
)

// NewGoodsCategoryModel returns a model for the database table.
func NewGoodsCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) GoodsCategoryModel {
	//用于插入数据的字段
	goodsCategoryRowsExpectAutoSet = strings.Join(stringx.Remove(goodsCategoryFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	goodsCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsCategoryFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customGoodsCategoryModel{
		defaultGoodsCategoryModel: newGoodsCategoryModel(conn, c),
	}
}

func (m *customGoodsCategoryModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheGoodsCategoryUuidPrefix, uuid)
}

// 软删除
func (m *customGoodsCategoryModel) DeleteSoft(ctx context.Context, data *GoodsCategory) error {
	goodsCategoryIdKey := m.formatPrimary(data.Id)
	goodsCategoryUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, goodsCategoryIdKey, goodsCategoryUuidKey)
	//完成
	return err
}

func (m *customGoodsCategoryModel) FindOneByUuid(ctx context.Context, uuid string) (*GoodsCategory, error) {
	goodsCategoryUuidKey := m.formatUuidKey(uuid)
	var resp GoodsCategory
	err := m.QueryRowCtx(ctx, &resp, goodsCategoryUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", goodsCategoryRows, m.table)
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

func (m *customGoodsCategoryModel) FindOne(ctx context.Context, id int64) (*GoodsCategory, error) {
	goodsCategoryIdKey := m.formatPrimary(id)
	var resp GoodsCategory
	err := m.QueryRowCtx(ctx, &resp, goodsCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ?  and `is_deleted` = '0' limit 1", goodsCategoryRows, m.table)
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

func (m *customGoodsCategoryModel) Insert(ctx context.Context, data *GoodsCategory) (sql.Result, error) {
	goodsCategoryIdKey := m.formatPrimary(data.Id)
	goodsCategoryUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, goodsCategoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Status)
	}, goodsCategoryIdKey, goodsCategoryUuidKey)
	return ret, err
}

func (m *customGoodsCategoryModel) Update(ctx context.Context, data *GoodsCategory) error {
	goodsCategoryIdKey := m.formatPrimary(data.Id)
	goodsCategoryUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, goodsCategoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Status, data.Id)
	}, goodsCategoryIdKey, goodsCategoryUuidKey)
	return err
}

//----------------------------------------------------------------

func (m *customGoodsCategoryModel) FindList(ctx context.Context, pageSize, page int64, keyword string, goodsCategory *GoodsCategory) (resp []*GoodsCategory, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(goodsCategoryRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if goodsCategory != nil {
		//name的优先级大于keyword
		if goodsCategory.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", goodsCategory.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", goodsCategory.Name, "%"))
			mark = true
		}
		if goodsCategory.Status != -1 {
			sq = sq.Where("`status` = ?", goodsCategory.Status)
			sqCount = sqCount.Where("`status` = ?", goodsCategory.Status)
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
