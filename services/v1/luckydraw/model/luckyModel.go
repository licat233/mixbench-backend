package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	_ LuckyModel = (*customLuckyModel)(nil)
)

type (
	// LuckyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLuckyModel.
	LuckyModel interface {
		luckyModel  // extended interface by gozero
		lucky_model // extended interface by genzero

		FindCountByUserId(ctx context.Context, userId int64) (int64, error)
		UpdateStatus(ctx context.Context, id, status int64) error
		UpdateSomeStatusById(ctx context.Context, ids []int64, status int64) error
		UpdateSomeStatusByUuid(ctx context.Context, uuidList []string, status int64) error
		// FindDetails(ctx context.Context, options *FindDetailsOptions) (resp []*LuckyDetail, total int64, err error)
		FindByUserAndActivity(ctx context.Context, UserId, ActivityId int64) ([]*Lucky, error)
		FindCountByUserAndActivity(ctx context.Context, UserId, ActivityId int64) (int64, error)
		// FindCountByPrizePoolId(ctx context.Context, prizesPoolIds []int64) (int64, error)
		FindlistByPrizePoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId, pageSize, page int64) (list []*Lucky, total int64, err error)
		// FindCountByPrizePoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId int64) (int64, error)

		FindlistByUserIds(ctx context.Context, userIds []int64, pageSize, page int64) (list []*Lucky, total int64, err error)
		FindlistByPrizePoolIds(ctx context.Context, prizesPoolIds []int64, pageSize, page int64) (list []*Lucky, total int64, err error)
		// FindCountByUserIds(ctx context.Context, userIds []int64) (int64, error)

		Conn() sqlc.CachedConn
	}

	customLuckyModel struct {
		*defaultLuckyModel
	}
)

// NewLuckyModel returns a model for the database table.
func NewLuckyModel(conn sqlx.SqlConn, c cache.CacheConf) LuckyModel {
	//用于插入数据的字段
	luckyRowsExpectAutoSet = strings.Join(stringx.Remove(luckyFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段，只有创建的时候才能设置ip
	luckyRowsWithPlaceHolder = strings.Join(stringx.Remove(luckyFieldNames, append(BaseUpdateOmitField, "ip")...), "=?,") + "=?"
	return &customLuckyModel{
		defaultLuckyModel: newLuckyModel(conn, c),
	}
}

func (m *customLuckyModel) FindCountByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where `user_id` = ? ", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, userId)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return count, ErrNotFound
	default:
		return count, err
	}
}

func (m *customLuckyModel) FindlistByPrizePoolIds(ctx context.Context, prizesPoolIds []int64, pageSize, page int64) (list []*Lucky, total int64, err error) {
	if len(prizesPoolIds) == 0 {
		return
	}
	query := fmt.Sprintf("select %s from %s where `prize_pool_id` in(?)", luckyRows, m.table)
	if pageSize > 0 && page > 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
		var count int64
		queryCount := fmt.Sprintf("select count(*) from %s where `prize_pool_id` in(?)", m.table)
		err = m.QueryRowNoCacheCtx(ctx, &count, queryCount, prizesPoolIds)
		total = count
		if err != nil {
			return
		}
	}
	err = m.QueryRowsNoCacheCtx(ctx, &list, query, prizesPoolIds)
	return
}

// func (m *customLuckyModel) FindCountByPrizePoolId(ctx context.Context, prizesPoolIds []int64) (int64, error) {
// 	if len(prizesPoolIds) == 0 {
// 		return 0, nil
// 	}
// 	var count int64
// 	var temps []string
// 	for _, prizesPoolId := range prizesPoolIds {
// 		temps = append(temps, fmt.Sprint(prizesPoolId))
// 	}
// 	whereSql := strings.Join(temps, ",")
// 	query := fmt.Sprintf("select count(*) from %s where `prize_pool_id` in(%s)", m.table, whereSql)
// 	err := m.QueryRowNoCacheCtx(ctx, &count, query)
// 	switch err {
// 	case nil:
// 		return count, nil
// 	case sqlc.ErrNotFound:
// 		return count, ErrNotFound
// 	default:
// 		return count, err
// 	}
// }

func (m *customLuckyModel) FindlistByPrizePoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId, pageSize, page int64) (list []*Lucky, total int64, err error) {
	if len(prizesPoolIds) == 0 || userId == 0 {
		return
	}
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `prize_pool_id` in(?)", luckyRows, m.table)
	if pageSize > 0 && page > 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
		var count int64
		queryCount := fmt.Sprintf("select count(*) from %s where  `user_id` = ? and `prize_pool_id` in(?)", m.table)
		err = m.QueryRowNoCacheCtx(ctx, &count, queryCount, userId, prizesPoolIds)
		total = count
		if err != nil {
			return
		}
	}
	err = m.QueryRowsNoCacheCtx(ctx, &list, query, userId, prizesPoolIds)
	return
}

// func (m *customLuckyModel) FindCountByPrizePoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId int64) (int64, error) {
// 	var count int64
// 	var temps []string
// 	if len(prizesPoolIds) == 0 {
// 		return 0, nil
// 	}
// 	for _, prizesPoolId := range prizesPoolIds {
// 		temps = append(temps, fmt.Sprint(prizesPoolId))
// 	}
// 	whereSql := strings.Join(temps, ",")
// 	query := fmt.Sprintf("select count(*) from %s where `user_id` = ? and `prize_pool_id` in(%s)", m.table, whereSql)
// 	err := m.QueryRowNoCacheCtx(ctx, &count, query, userId)
// 	return count, err
// }

func (m *customLuckyModel) UpdateStatus(ctx context.Context, id, status int64) error {
	// luckyIdKey := fmt.Sprintf("%s%v", cacheLuckyIdPrefix, id)
	luckyIdKey := m.formatPrimary(id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ? where `id` = ? ", m.table)
		return conn.ExecCtx(ctx, query, status, id)
	}, luckyIdKey)
	return err
}

func (m *customLuckyModel) UpdateSomeStatusById(ctx context.Context, ids []int64, status int64) error {
	var keys []string
	var temps []string
	args := []interface{}{status}
	for _, id := range ids {
		temps = append(temps, "?")
		luckyIdKey := m.formatPrimary(id)
		keys = append(keys, luckyIdKey)
		args = append(args, id)
	}
	temp := strings.Join(temps, ",")

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ? where `id` in (%s) ", m.table, temp)
		return conn.ExecCtx(ctx, query, args...)
	}, keys...)
	return err
}

func (m *customLuckyModel) UpdateSomeStatusByUuid(ctx context.Context, uuidList []string, status int64) error {
	var temps []string
	args := []interface{}{status}
	for _, uuid := range uuidList {
		temps = append(temps, "?")
		args = append(args, uuid)
	}
	temp := strings.Join(temps, ",")
	query := fmt.Sprintf("update %s set `status`= ? where `uuid` in (%s) ", m.table, temp)
	_, err := m.ExecNoCacheCtx(ctx, query, args...)
	return err
}

func (m *customLuckyModel) FindByUserAndActivity(ctx context.Context, UserId, ActivityId int64) ([]*Lucky, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? and activity_id = ? ", luckyRows, m.table)
	var resp []*Lucky
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, UserId, ActivityId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckyModel) FindCountByUserAndActivity(ctx context.Context, UserId, ActivityId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where user_id = ? and activity_id = ? ", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, UserId, ActivityId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customLuckyModel) FindlistByUserIds(ctx context.Context, userIds []int64, pageSize, page int64) (list []*Lucky, total int64, err error) {
	if len(userIds) == 0 {
		return
	}
	query := fmt.Sprintf("select %s from %s where `user_id` in(?)", luckyRows, m.table)
	if pageSize > 0 && page > 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
		var count int64
		queryCount := fmt.Sprintf("select count(*) from %s where `user_id` in(?)", m.table)
		err = m.QueryRowNoCacheCtx(ctx, &count, queryCount, userIds)
		total = count
		if err != nil {
			return
		}
	}
	err = m.QueryRowsNoCacheCtx(ctx, &list, query, userIds)
	return
}

// func (m *customLuckyModel) FindCountByUserIds(ctx context.Context, userIds []int64) (int64, error) {
// 	var count int64
// 	var temps []string
// 	if len(userIds) == 0 {
// 		return 0, nil
// 	}
// 	for _, userId := range userIds {
// 		temps = append(temps, fmt.Sprint(userId))
// 	}
// 	whereSql := strings.Join(temps, ",")
// 	query := fmt.Sprintf("select count(*) from %s where `user_id` in(%s)", m.table, whereSql)
// 	err := m.QueryRowsNoCacheCtx(ctx, &count, query)
// 	return count, err
// }

func (m *customLuckyModel) Conn() sqlc.CachedConn {
	return m.CachedConn
}
