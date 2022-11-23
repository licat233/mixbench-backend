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

var _ PrizeGradeModel = (*customPrizeGradeModel)(nil)
var PrizeGradeTableName = "prize_grade"

type (
	// PrizeGradeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeGradeModel.
	PrizeGradeModel interface {
		prizeGradeModel
		DeleteSoft(ctx context.Context, data *PrizeGrade) error

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, prizeGrade *PrizeGrade) (resp []*PrizeGrade, total int64, err error)
		TableName() string
	}

	customPrizeGradeModel struct {
		*defaultPrizeGradeModel
	}
)

// NewPrizeGradeModel returns a model for the database table.
func NewPrizeGradeModel(conn sqlx.SqlConn, c cache.CacheConf) PrizeGradeModel {
	//用于插入数据的字段
	prizeGradeRowsExpectAutoSet = strings.Join(stringx.Remove(prizeGradeFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	prizeGradeRowsWithPlaceHolder = strings.Join(stringx.Remove(prizeGradeFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customPrizeGradeModel{
		defaultPrizeGradeModel: newPrizeGradeModel(conn, c),
	}
}

func (m *customPrizeGradeModel) Delete(ctx context.Context, id int64) error {
	luckDrawPrizeGradeIdKey := m.formatPrimary(id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, luckDrawPrizeGradeIdKey)
	return err
}

// 软删除
func (m *customPrizeGradeModel) DeleteSoft(ctx context.Context, data *PrizeGrade) error {
	luckDrawPrizeGradeIdKey := m.formatPrimary(data.Id)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawPrizeGradeIdKey)
	//完成
	return err
}

func (m *customPrizeGradeModel) FindOne(ctx context.Context, id int64) (*PrizeGrade, error) {
	luckDrawPrizeGradeIdKey := m.formatPrimary(id)
	var resp PrizeGrade
	err := m.QueryRowCtx(ctx, &resp, luckDrawPrizeGradeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' limit 1", prizeGradeRows, m.table)
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

func (m *customPrizeGradeModel) Insert(ctx context.Context, data *PrizeGrade) (sql.Result, error) {
	luckDrawPrizeGradeIdKey := m.formatPrimary(data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, prizeGradeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.IsWin)
	}, luckDrawPrizeGradeIdKey)
	return ret, err
}

func (m *customPrizeGradeModel) Update(ctx context.Context, data *PrizeGrade) error {
	luckDrawPrizeGradeIdKey := m.formatPrimary(data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, prizeGradeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.IsWin, data.Id)
	}, luckDrawPrizeGradeIdKey)
	return err
}

//----------------------------------------------------------------

func (m *customPrizeGradeModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `is_deleted`= '0'", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *customPrizeGradeModel) FindList(ctx context.Context, pageSize, page int64, keyword string, prizeGrade *PrizeGrade) (resp []*PrizeGrade, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(prizeGradeRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if prizeGrade != nil {
		//name的优先级大于keyword
		if prizeGrade.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", prizeGrade.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", prizeGrade.Name, "%"))
			mark = true
		}
		if prizeGrade.IsWin != -1 {
			sq = sq.Where("`is_win` = ?", prizeGrade.IsWin)
			sqCount = sqCount.Where("`is_win` = ?", prizeGrade.IsWin)
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

func (m *customPrizeGradeModel) TableName() string {
	return m.table
}
