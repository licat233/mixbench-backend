/*
 * @Author: licat
 * @Date: 2023-01-10 15:16:02
 * @LastEditors: licat
 * @LastEditTime: 2023-01-11 15:00:48
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

var _ GoodsModel = (*customGoodsModel)(nil)
var GoodsTableName = "goods"
var cacheGoodsUuidPrefix = "cache:goods:uuid:"

type (
	// GoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsModel.
	GoodsModel interface {
		goodsModel
		DeleteSoft(ctx context.Context, data *Goods) error
		FindOneByUuid(ctx context.Context, uuid string) (*Goods, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, goods *Goods) (resp []*Goods, total int64, err error)
	}

	customGoodsModel struct {
		*defaultGoodsModel
	}
)

// NewGoodsModel returns a model for the database table.
func NewGoodsModel(conn sqlx.SqlConn, c cache.CacheConf) GoodsModel {
	//用于插入数据的字段
	goodsRowsExpectAutoSet = strings.Join(stringx.Remove(goodsFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	goodsRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customGoodsModel{
		defaultGoodsModel: newGoodsModel(conn, c),
	}
}

func (m *customGoodsModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheGoodsUuidPrefix, uuid)
}

// 软删除
func (m *customGoodsModel) DeleteSoft(ctx context.Context, data *Goods) error {
	goodsIdKey := m.formatPrimary(data.Id)
	goodsUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, goodsIdKey, goodsUuidKey)
	//完成
	return err
}

func (m *customGoodsModel) FindOneByUuid(ctx context.Context, uuid string) (*Goods, error) {
	goodsUuidKey := m.formatUuidKey(uuid)
	var resp Goods
	err := m.QueryRowCtx(ctx, &resp, goodsUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", goodsRows, m.table)
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

func (m *customGoodsModel) FindOne(ctx context.Context, id int64) (*Goods, error) {
	goodsIdKey := m.formatPrimary(id)
	var resp Goods
	err := m.QueryRowCtx(ctx, &resp, goodsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ?  and `is_deleted` = '0' limit 1", goodsRows, m.table)
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

func (m *customGoodsModel) Insert(ctx context.Context, data *Goods) (sql.Result, error) {
	goodsIdKey := m.formatPrimary(data.Id)
	goodsUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, goodsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Status)
	}, goodsIdKey, goodsUuidKey)
	return ret, err
}

func (m *customGoodsModel) Update(ctx context.Context, data *Goods) error {
	goodsIdKey := m.formatPrimary(data.Id)
	goodsUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, goodsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Status, data.Id)
	}, goodsIdKey, goodsUuidKey)
	return err
}

//----------------------------------------------------------------

func (m *customGoodsModel) FindList(ctx context.Context, pageSize, page int64, keyword string, goods *Goods) (resp []*Goods, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(goodsRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if goods != nil {
		//name的优先级大于keyword
		if goods.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", goods.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", goods.Name, "%"))
			mark = true
		}
		if goods.Status != -1 {
			sq = sq.Where("`status` = ?", goods.Status)
			sqCount = sqCount.Where("`status` = ?", goods.Status)
			mark = true
		}
		if goods.IsReal != -1 {
			sq = sq.Where("`is_real` = ?", goods.IsReal)
			sqCount = sqCount.Where("`is_real` = ?", goods.IsReal)
			mark = true
		}
		if goods.IsOnSale != -1 {
			sq = sq.Where("`is_on_sale` = ?", goods.IsOnSale)
			sqCount = sqCount.Where("`is_on_sale` = ?", goods.IsOnSale)
			mark = true
		}
		if goods.IsAloneSale != -1 {
			sq = sq.Where("`is_alone_sale` = ?", goods.IsAloneSale)
			sqCount = sqCount.Where("`is_alone_sale` = ?", goods.IsAloneSale)
			mark = true
		}
		if goods.IsBest != -1 {
			sq = sq.Where("`is_best` = ?", goods.IsBest)
			sqCount = sqCount.Where("`is_best` = ?", goods.IsBest)
			mark = true
		}
		if goods.IsNew != -1 {
			sq = sq.Where("`is_new` = ?", goods.IsNew)
			sqCount = sqCount.Where("`is_new` = ?", goods.IsNew)
			mark = true
		}
		if goods.IsHot != -1 {
			sq = sq.Where("`is_hot` = ?", goods.IsHot)
			sqCount = sqCount.Where("`is_hot` = ?", goods.IsHot)
			mark = true
		}
		if goods.Tag != "" {
			sq = sq.Where("`tag` = ?", goods.Tag)
			sqCount = sqCount.Where("`tag` = ?", goods.Tag)
			mark = true
		}
		if goods.Keywords != "" {
			sq = sq.Where("`keywords` = ?", goods.Keywords)
			sqCount = sqCount.Where("`keywords` = ?", goods.Keywords)
			mark = true
		}
		if goods.GoodsPrice != -1 {
			sq = sq.Where("`goods_price` = ?", goods.GoodsPrice)
			sqCount = sqCount.Where("`goods_price` = ?", goods.GoodsPrice)
			mark = true
		}
		if goods.RemainCount != -1 {
			sq = sq.Where("`remain_count` = ?", goods.RemainCount)
			sqCount = sqCount.Where("`remain_count` = ?", goods.RemainCount)
			mark = true
		}
		if goods.MonthSellCount != -1 {
			sq = sq.Where("`month_sell_count` = ?", goods.MonthSellCount)
			sqCount = sqCount.Where("`month_sell_count` = ?", goods.MonthSellCount)
			mark = true
		}
		if goods.TotalSellCount != -1 {
			sq = sq.Where("`total_sell_count` = ?", goods.TotalSellCount)
			sqCount = sqCount.Where("`total_sell_count` = ?", goods.TotalSellCount)
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
