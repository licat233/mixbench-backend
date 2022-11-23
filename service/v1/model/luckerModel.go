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
	_                             LuckerModel = (*customLuckerModel)(nil)
	LuckerTableName                           = "lucker"
	cacheLuckDrawLuckerUuidPrefix             = "cache:luckDraw:lucker:uuid:"
)

type (
	// LuckerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLuckerModel.
	LuckerModel interface {
		luckerModel
		DeleteSoft(ctx context.Context, data *Lucker) error
		FindOneByUuid(ctx context.Context, uuid string) (*Lucker, error)

		FindList(ctx context.Context, pageSize, page int64, keyword string, lucker *Lucker) (resp []*Lucker, total int64, err error)
		FindsByUserId(ctx context.Context, userId int64, pageSize, page int64) ([]*Lucker, error)
		FindCountByUserId(ctx context.Context, userId int64) (int64, error)
		FindCount(ctx context.Context) (int64, error)
		UpdateStatus(ctx context.Context, id, status int64) error
		UpdateSomeStatusById(ctx context.Context, ids []int64, status int64) error
		UpdateSomeStatusByUuid(ctx context.Context, uuidList []string, status int64) error
		// FindDetails(ctx context.Context, options *FindDetailsOptions) (resp []*LuckerDetail, total int64, err error)
		FindByUserAndActivity(ctx context.Context, UserId, ActivityId int64) ([]*Lucker, error)
		FindCountByUserAndActivity(ctx context.Context, UserId, ActivityId int64) (int64, error)
		FindsByPrizesPoolId(ctx context.Context, prizesPoolIds []int64, pageSize, page int64) ([]*Lucker, error)
		FindCountByPrizesPoolId(ctx context.Context, prizesPoolIds []int64) (int64, error)
		FindsByPrizesPoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId, pageSize, page int64) ([]*Lucker, error)
		FindCountByPrizesPoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId int64) (int64, error)

		FindsByUserIds(ctx context.Context, userIds []int64, pageSize, page int64) ([]*Lucker, error)
		FindCountByUserIds(ctx context.Context, userIds []int64) (int64, error)

		TableName() string
		Conn() sqlc.CachedConn
	}

	customLuckerModel struct {
		*defaultLuckerModel
	}
)

// NewLuckerModel returns a model for the database table.
func NewLuckerModel(conn sqlx.SqlConn, c cache.CacheConf) LuckerModel {
	//用于插入数据的字段
	luckerRowsExpectAutoSet = strings.Join(stringx.Remove(luckerFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段，只有创建的时候才能设置ip
	luckerRowsWithPlaceHolder = strings.Join(stringx.Remove(luckerFieldNames, append(BaseUpdateOmitField, "ip")...), "=?,") + "=?"
	return &customLuckerModel{
		defaultLuckerModel: newLuckerModel(conn, c),
	}
}

func (m *customLuckerModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawLuckerUuidPrefix, uuid)
}

// 软删除
func (m *customLuckerModel) DeleteSoft(ctx context.Context, data *Lucker) error {
	luckDrawLuckerIdKey := m.formatPrimary(data.Id)
	luckDrawLuckerUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawLuckerIdKey, luckDrawLuckerUuidKey)
	//完成
	return err
}

func (m *customLuckerModel) FindOneByUuid(ctx context.Context, uuid string) (*Lucker, error) {
	luckDrawLuckerUuidKey := m.formatUuidKey(uuid)
	var resp Lucker
	err := m.QueryRowCtx(ctx, &resp, luckDrawLuckerUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", luckerRows, m.table)
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

func (m *customLuckerModel) FindOne(ctx context.Context, id int64) (*Lucker, error) {
	luckDrawLuckerIdKey := fmt.Sprintf("%s%v", cacheLuckDrawLuckerIdPrefix, id)
	var resp Lucker
	err := m.QueryRowCtx(ctx, &resp, luckDrawLuckerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' limit 1", luckerRows, m.table)
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

func (m *customLuckerModel) Insert(ctx context.Context, data *Lucker) (sql.Result, error) {
	if data.Ip == "" {
		data.Ip = "0.0.0.0"
	}
	luckDrawLuckerIdKey := m.formatPrimary(data.Id)
	luckDrawLuckerUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, luckerRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.UserId, data.PrizesPoolId, data.Status, data.LuckDate, data.Ip)
	}, luckDrawLuckerIdKey, luckDrawLuckerUuidKey)
	return ret, err
}

func (m *customLuckerModel) Update(ctx context.Context, data *Lucker) error {
	luckDrawLuckerIdKey := m.formatPrimary(data.Id)
	luckDrawLuckerUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, luckerRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.PrizesPoolId, data.Status, data.LuckDate, data.Ip, data.Id)
	}, luckDrawLuckerIdKey, luckDrawLuckerUuidKey)
	return err
}

//----------------------------------------------------------------------------------------------------------------------

func (m *customLuckerModel) FindCount(ctx context.Context) (int64, error) {
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

func (m *customLuckerModel) FindList(ctx context.Context, pageSize, page int64, keyword string, lucker *Lucker) (resp []*Lucker, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(luckerRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	if lucker != nil {
		if utils.IdValid(lucker.UserId) {
			sq = sq.Where("`user_id` = ?", lucker.UserId)
			sqCount = sqCount.Where("`user_id` = ?", lucker.UserId)
		}
		if utils.IdValid(lucker.PrizesPoolId) {
			sq = sq.Where("`prizes_pool_id` = ?", lucker.PrizesPoolId)
			sqCount = sqCount.Where("`prizes_pool_id` = ?", lucker.PrizesPoolId)
		}
		if lucker.Status != -1 {
			sq = sq.Where("`status` = ?", lucker.Status)
			sqCount = sqCount.Where("`status` = ?", lucker.Status)
		}
		if !lucker.LuckDate.IsZero() {
			dayStr := lucker.LuckDate.Format("2006-01-02")
			sq = sq.Where("DATE_FORMAT(`luck_date`,'%Y-%m-%d') = ?", dayStr)
			sqCount = sqCount.Where("DATE_FORMAT(`luck_date`,'%Y-%m-%d') = ?", dayStr)
		}
		if lucker.Ip != "" {
			sq = sq.Where("`ip` = ?", lucker.Ip)
			sqCount = sqCount.Where("`ip` = ?", lucker.Ip)
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

func (m *customLuckerModel) FindsByUserId(ctx context.Context, userId int64, pageSize, page int64) ([]*Lucker, error) {
	var resp []*Lucker
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `is_deleted` = '0'", luckerRows, m.table)
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

func (m *customLuckerModel) FindCountByUserId(ctx context.Context, userId int64) (int64, error) {
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

func (m *customLuckerModel) FindsByPrizesPoolId(ctx context.Context, prizesPoolIds []int64, pageSize, page int64) ([]*Lucker, error) {
	if len(prizesPoolIds) == 0 {
		return nil, nil
	}
	var resp []*Lucker
	var temps []string
	for _, prizesPoolId := range prizesPoolIds {
		temps = append(temps, fmt.Sprint(prizesPoolId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select %s from %s where `is_deleted` = '0' and `prizes_pool_id` in(%s)", luckerRows, m.table, whereSql)
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

func (m *customLuckerModel) FindCountByPrizesPoolId(ctx context.Context, prizesPoolIds []int64) (int64, error) {
	if len(prizesPoolIds) == 0 {
		return 0, nil
	}
	var count int64
	var temps []string
	for _, prizesPoolId := range prizesPoolIds {
		temps = append(temps, fmt.Sprint(prizesPoolId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select count(*) from %s where `is_deleted` = '0' and `prizes_pool_id` in(%s)", m.table, whereSql)
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

func (m *customLuckerModel) FindsByPrizesPoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId, pageSize, page int64) ([]*Lucker, error) {
	var resp []*Lucker
	var temps []string
	if len(prizesPoolIds) == 0 {
		return nil, nil
	}
	for _, prizesPoolId := range prizesPoolIds {
		temps = append(temps, fmt.Sprint(prizesPoolId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select %s from %s where `is_deleted` = '0' and `user_id` = ? and `prizes_pool_id` in(%s)", luckerRows, m.table, whereSql)
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

func (m *customLuckerModel) FindCountByPrizesPoolIdAndUserId(ctx context.Context, prizesPoolIds []int64, userId int64) (int64, error) {
	var count int64
	var temps []string
	if len(prizesPoolIds) == 0 {
		return 0, nil
	}
	for _, prizesPoolId := range prizesPoolIds {
		temps = append(temps, fmt.Sprint(prizesPoolId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select count(*) from %s where `is_deleted` = '0' and `user_id` = ? and `prizes_pool_id` in(%s)", m.table, whereSql)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, userId)
	return count, err
}

func (m *customLuckerModel) UpdateStatus(ctx context.Context, id, status int64) error {
	luckDrawLuckerIdKey := fmt.Sprintf("%s%v", cacheLuckDrawLuckerIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ? where `id` = ? and `is_deleted` = '0'", m.table)
		return conn.ExecCtx(ctx, query, status, id)
	}, luckDrawLuckerIdKey)
	return err
}

func (m *customLuckerModel) UpdateSomeStatusById(ctx context.Context, ids []int64, status int64) error {
	var keys []string
	var temps []string
	args := []interface{}{status}
	for _, id := range ids {
		temps = append(temps, "?")
		luckDrawLuckerIdKey := fmt.Sprintf("%s%v", cacheLuckDrawLuckerIdPrefix, id)
		keys = append(keys, luckDrawLuckerIdKey)
		args = append(args, id)
	}
	temp := strings.Join(temps, ",")

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ? where `id` in (%s) and `is_deleted` = '0'", m.table, temp)
		return conn.ExecCtx(ctx, query, args...)
	}, keys...)
	return err
}

func (m *customLuckerModel) UpdateSomeStatusByUuid(ctx context.Context, uuidList []string, status int64) error {
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

func (m *customLuckerModel) FindByUserAndActivity(ctx context.Context, UserId, ActivityId int64) ([]*Lucker, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? and activity_id = ? and is_deleted = '0'", luckerRows, m.table)
	var resp []*Lucker
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

func (m *customLuckerModel) FindCountByUserAndActivity(ctx context.Context, UserId, ActivityId int64) (int64, error) {
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

func (m *customLuckerModel) FindsByUserIds(ctx context.Context, userIds []int64, pageSize, page int64) ([]*Lucker, error) {
	var resp []*Lucker
	var temps []string
	if len(userIds) == 0 {
		return nil, nil
	}
	for _, userId := range userIds {
		temps = append(temps, fmt.Sprint(userId))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select %s from %s where `is_deleted` = '0' and `user_id` in(%s)", luckerRows, m.table, whereSql)
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

func (m *customLuckerModel) FindCountByUserIds(ctx context.Context, userIds []int64) (int64, error) {
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

func (m *customLuckerModel) TableName() string {
	return m.table
}

func (m *customLuckerModel) Conn() sqlc.CachedConn {
	return m.CachedConn
}
