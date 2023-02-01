package model

import (
	"context"
	"database/sql"
	"fmt"
	"mixbench/common/utils"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	_                    LuckyModel = (*customLuckyModel)(nil)
	LuckyTableName                  = "lucky"
	cacheLuckyUuidPrefix            = "cache:lucky:uuid:"
)

type (
	// LuckyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLuckyModel.
	LuckyModel interface {
		luckyModel
		DeleteSoft(ctx context.Context, data *Lucky) error
		FindOneByUuid(ctx context.Context, uuid string) (*Lucky, error)

		FindList(ctx context.Context, pageSize, page int64, keyword string, lucky *Lucky) (resp []*Lucky, total int64, err error)
		FindsByUserId(ctx context.Context, userId int64, pageSize, page int64) ([]*Lucky, error)
		FindCountByUserId(ctx context.Context, userId int64) (int64, error)
		FindCount(ctx context.Context) (int64, error)
		UpdateStatus(ctx context.Context, id, status int64) error
		UpdateSomeStatusById(ctx context.Context, ids []int64, status int64) error
		UpdateSomeStatusByUuid(ctx context.Context, uuidList []string, status int64) error
		// FindDetails(ctx context.Context, options *FindDetailsOptions) (resp []*LuckyDetail, total int64, err error)
		FindByUserAndActivity(ctx context.Context, UserId, ActivityId int64) ([]*Lucky, error)
		FindCountByUserAndActivity(ctx context.Context, UserId, ActivityId int64) (int64, error)
		FindsByPrizePoolId(ctx context.Context, prizesPoolIds []int64, pageSize, page int64) ([]*Lucky, error)
		FindCountByPrizePoolId(ctx context.Context, prizesPoolIds []int64) (int64, error)
		FindsByPrizePoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId, pageSize, page int64) ([]*Lucky, error)
		FindCountByPrizePoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId int64) (int64, error)

		FindsByUserIds(ctx context.Context, userIds []int64, pageSize, page int64) ([]*Lucky, error)
		FindCountByUserIds(ctx context.Context, userIds []int64) (int64, error)

		TableName() string
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

func (m *customLuckyModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckyUuidPrefix, uuid)
}

// 软删除
func (m *customLuckyModel) DeleteSoft(ctx context.Context, data *Lucky) error {
	luckyIdKey := m.formatPrimary(data.Id)
	luckyUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckyIdKey, luckyUuidKey)
	//完成
	return err
}

func (m *customLuckyModel) FindOneByUuid(ctx context.Context, uuid string) (*Lucky, error) {
	luckyUuidKey := m.formatUuidKey(uuid)
	var resp Lucky
	err := m.QueryRowCtx(ctx, &resp, luckyUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", luckyRows, m.table)
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

func (m *customLuckyModel) FindOne(ctx context.Context, id int64) (*Lucky, error) {
	luckyIdKey := m.formatPrimary(id)
	var resp Lucky
	err := m.QueryRowCtx(ctx, &resp, luckyIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' limit 1", luckyRows, m.table)
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

func (m *customLuckyModel) Insert(ctx context.Context, data *Lucky) (sql.Result, error) {
	if data.Ip == "" {
		data.Ip = "0.0.0.0"
	}
	luckyIdKey := m.formatPrimary(data.Id)
	luckyUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, luckyRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.UserId, data.PrizePoolId, data.Status, data.LuckDate, data.Ip)
	}, luckyIdKey, luckyUuidKey)
	return ret, err
}

func (m *customLuckyModel) Update(ctx context.Context, data *Lucky) error {
	luckyIdKey := m.formatPrimary(data.Id)
	luckyUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, luckyRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.PrizePoolId, data.Status, data.LuckDate, data.Ip, data.Id)
	}, luckyIdKey, luckyUuidKey)
	return err
}

//----------------------------------------------------------------------------------------------------------------------

func (m *customLuckyModel) FindCount(ctx context.Context) (int64, error) {
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

func (m *customLuckyModel) FindList(ctx context.Context, pageSize, page int64, keyword string, lucky *Lucky) (resp []*Lucky, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(luckyRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	if lucky != nil {
		if utils.IdValid(lucky.UserId) {
			sq = sq.Where("`user_id` = ?", lucky.UserId)
			sqCount = sqCount.Where("`user_id` = ?", lucky.UserId)
		}
		if utils.IdValid(lucky.PrizePoolId) {
			sq = sq.Where("`prize_pool_id` = ?", lucky.PrizePoolId)
			sqCount = sqCount.Where("`prize_pool_id` = ?", lucky.PrizePoolId)
		}
		if lucky.Status != -1 {
			sq = sq.Where("`status` = ?", lucky.Status)
			sqCount = sqCount.Where("`status` = ?", lucky.Status)
		}
		if !lucky.LuckDate.IsZero() {
			dayStr := lucky.LuckDate.Format("2006-01-02")
			sq = sq.Where("DATE_FORMAT(`luck_date`,'%Y-%m-%d') = ?", dayStr)
			sqCount = sqCount.Where("DATE_FORMAT(`luck_date`,'%Y-%m-%d') = ?", dayStr)
		}
		if lucky.Ip != "" {
			sq = sq.Where("`ip` = ?", lucky.Ip)
			sqCount = sqCount.Where("`ip` = ?", lucky.Ip)
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

func (m *customLuckyModel) FindsByUserId(ctx context.Context, userId int64, pageSize, page int64) ([]*Lucky, error) {
	var resp []*Lucky
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `is_deleted` = '0'", luckyRows, m.table)
	if pageSize > 0 && page > 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
	}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckyModel) FindCountByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where `user_id` = ? and `is_deleted` = '0'", m.table)
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

func (m *customLuckyModel) FindsByPrizePoolId(ctx context.Context, prizesPoolIds []int64, pageSize, page int64) ([]*Lucky, error) {
	if len(prizesPoolIds) == 0 {
		return nil, nil
	}
	var resp []*Lucky
	var temps []string
	for _, prizesPoolId := range prizesPoolIds {
		temps = append(temps, fmt.Sprint(prizesPoolId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select %s from %s where `is_deleted` = '0' and `prize_pool_id` in(%s)", luckyRows, m.table, whereSql)
	if pageSize > 0 && page > 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
	}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckyModel) FindCountByPrizePoolId(ctx context.Context, prizesPoolIds []int64) (int64, error) {
	if len(prizesPoolIds) == 0 {
		return 0, nil
	}
	var count int64
	var temps []string
	for _, prizesPoolId := range prizesPoolIds {
		temps = append(temps, fmt.Sprint(prizesPoolId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select count(*) from %s where `is_deleted` = '0' and `prize_pool_id` in(%s)", m.table, whereSql)
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return count, ErrNotFound
	default:
		return count, err
	}
}

func (m *customLuckyModel) FindsByPrizePoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId, pageSize, page int64) ([]*Lucky, error) {
	var resp []*Lucky
	var temps []string
	if len(prizesPoolIds) == 0 {
		return nil, nil
	}
	for _, prizesPoolId := range prizesPoolIds {
		temps = append(temps, fmt.Sprint(prizesPoolId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select %s from %s where `is_deleted` = '0' and `user_id` = ? and `prize_pool_id` in(%s)", luckyRows, m.table, whereSql)
	if pageSize > 0 && page > 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
	}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckyModel) FindCountByPrizePoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId int64) (int64, error) {
	var count int64
	var temps []string
	if len(prizesPoolIds) == 0 {
		return 0, nil
	}
	for _, prizesPoolId := range prizesPoolIds {
		temps = append(temps, fmt.Sprint(prizesPoolId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select count(*) from %s where `is_deleted` = '0' and `user_id` = ? and `prize_pool_id` in(%s)", m.table, whereSql)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, userId)
	return count, err
}

func (m *customLuckyModel) UpdateStatus(ctx context.Context, id, status int64) error {
	// luckyIdKey := fmt.Sprintf("%s%v", cacheLuckyIdPrefix, id)
	luckyIdKey := m.formatPrimary(id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ? where `id` = ? and `is_deleted` = '0'", m.table)
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
		query := fmt.Sprintf("update %s set `status`= ? where `id` in (%s) and `is_deleted` = '0'", m.table, temp)
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
	query := fmt.Sprintf("update %s set `status`= ? where `uuid` in (%s) and `is_deleted` = '0'", m.table, temp)
	_, err := m.ExecNoCacheCtx(ctx, query, args...)
	return err
}

func (m *customLuckyModel) FindByUserAndActivity(ctx context.Context, UserId, ActivityId int64) ([]*Lucky, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? and activity_id = ? and is_deleted = '0'", luckyRows, m.table)
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
	query := fmt.Sprintf("select count(*) from %s where user_id = ? and activity_id = ? and is_deleted = '0'", m.table)
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

func (m *customLuckyModel) FindsByUserIds(ctx context.Context, userIds []int64, pageSize, page int64) ([]*Lucky, error) {
	var resp []*Lucky
	var temps []string
	if len(userIds) == 0 {
		return nil, nil
	}
	for _, userId := range userIds {
		temps = append(temps, fmt.Sprint(userId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select %s from %s where `is_deleted` = '0' and `user_id` in(%s)", luckyRows, m.table, whereSql)
	if pageSize > 0 && page > 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
	}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckyModel) FindCountByUserIds(ctx context.Context, userIds []int64) (int64, error) {
	var count int64
	var temps []string
	if len(userIds) == 0 {
		return 0, nil
	}
	for _, userId := range userIds {
		temps = append(temps, fmt.Sprint(userId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select count(*) from %s where `is_deleted` = '0' and `user_id` in(%s)", m.table, whereSql)
	err := m.QueryRowsNoCacheCtx(ctx, &count, query)
	return count, err
}

func (m *customLuckyModel) TableName() string {
	return m.table
}

func (m *customLuckyModel) Conn() sqlc.CachedConn {
	return m.CachedConn
}
