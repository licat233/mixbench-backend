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

var _ UsersPoolModel = (*customUsersPoolModel)(nil)
var UsersPoolTableName = "users_pool"
var cacheLuckDrawUsersPoolUuidPrefix = "cache:luckDraw:usersPool:uuid:"

type (
	// UsersPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersPoolModel.
	UsersPoolModel interface {
		usersPoolModel
		DeleteSoft(ctx context.Context, data *UsersPool) error
		FindOneByUuid(ctx context.Context, uuid string) (*UsersPool, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64, usersPool *UsersPool) ([]*UsersPool, int64, error)
		FindsByUserId(ctx context.Context, userId int64) ([]*UsersPool, error)
		FindByUidAndAid(ctx context.Context, userId, activityId int64) ([]*UsersPool, error)
	}

	customUsersPoolModel struct {
		*defaultUsersPoolModel
	}
)

// NewUsersPoolModel returns a model for the database table.
func NewUsersPoolModel(conn sqlx.SqlConn, c cache.CacheConf) UsersPoolModel {
	//用于插入数据的字段
	usersPoolRowsExpectAutoSet = strings.Join(stringx.Remove(usersPoolFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	usersPoolRowsWithPlaceHolder = strings.Join(stringx.Remove(usersPoolFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customUsersPoolModel{
		defaultUsersPoolModel: newUsersPoolModel(conn, c),
	}
}

func (m *customUsersPoolModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawUsersPoolUuidPrefix, uuid)
}

// 软删除
func (m *customUsersPoolModel) DeleteSoft(ctx context.Context, data *UsersPool) error {
	luckDrawUsersPoolIdKey := m.formatPrimary(data.Id)
	luckDrawUsersPoolUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawUsersPoolIdKey, luckDrawUsersPoolUuidKey)
	//完成
	return err
}

func (m *customUsersPoolModel) FindOneByUuid(ctx context.Context, uuid string) (*UsersPool, error) {
	luckDrawUsersPoolUuidKey := m.formatUuidKey(uuid)
	var resp UsersPool
	err := m.QueryRowCtx(ctx, &resp, luckDrawUsersPoolUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", usersPoolRows, m.table)
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

func (m *customUsersPoolModel) FindOne(ctx context.Context, id int64) (*UsersPool, error) {
	luckDrawUsersPoolIdKey := m.formatPrimary(id)
	var resp UsersPool
	err := m.QueryRowCtx(ctx, &resp, luckDrawUsersPoolIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ?  and `is_deleted` = '0' limit 1", usersPoolRows, m.table)
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

func (m *customUsersPoolModel) Insert(ctx context.Context, data *UsersPool) (sql.Result, error) {
	luckDrawUsersPoolIdKey := m.formatPrimary(data.Id)
	luckDrawUsersPoolUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, usersPoolRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.ActivityId, data.UserId, data.Status, data.Score)
	}, luckDrawUsersPoolIdKey, luckDrawUsersPoolUuidKey)
	return ret, err
}

func (m *customUsersPoolModel) Update(ctx context.Context, data *UsersPool) error {
	luckDrawUsersPoolIdKey := m.formatPrimary(data.Id)
	luckDrawUsersPoolUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0' ", m.table, usersPoolRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ActivityId, data.UserId, data.Status, data.Score, data.Id)
	}, luckDrawUsersPoolIdKey, luckDrawUsersPoolUuidKey)
	return err
}

//----------------------------------------------------------------

func (m *customUsersPoolModel) FindCount(ctx context.Context) (int64, error) {
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

func (m *customUsersPoolModel) FindList(ctx context.Context, pageSize, page int64, usersPool *UsersPool) (resp []*UsersPool, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sqQuery := squirrel.Select(usersPoolRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	if usersPool != nil {
		if usersPool.ActivityId != -1 {
			sqQuery = sqQuery.Where("`activity_id` = ?", usersPool.ActivityId)
			sqCount = sqCount.Where("`activity_id` = ?", usersPool.ActivityId)
		}
		if usersPool.UserId != -1 {
			sqQuery = sqQuery.Where("`user_id` = ?", usersPool.UserId)
			sqCount = sqCount.Where("`user_id` = ?", usersPool.UserId)
		}
		if usersPool.Status != -1 {
			sqQuery = sqQuery.Where("`status` = ?", usersPool.Status)
			sqCount = sqCount.Where("`status` = ?", usersPool.Status)
		}
		if usersPool.Score != -1 {
			sqQuery = sqQuery.Where("`score` = ?", usersPool.Score)
			sqCount = sqCount.Where("`score` = ?", usersPool.Score)
		}
	}

	isPageQuery := pageSize > 0 && page > 0
	if isPageQuery {
		sqQuery = sqQuery.Offset(uint64((page - 1) * pageSize)).Limit(uint64(pageSize))
		queryCount, agrsCount, e := sqCount.ToSql()
		if e != nil {
			err = e
			return
		}
		if err = m.QueryRowNoCacheCtx(ctx, &total, queryCount, agrsCount...); err != nil {
			return
		}
	}

	query, agrs, err := sqQuery.ToSql()
	if err != nil {
		return
	}
	if err = m.QueryRowsNoCacheCtx(ctx, &resp, query, agrs...); err != nil {
		return
	}
	return
}

func (m *customUsersPoolModel) FindsByUserId(ctx context.Context, userId int64) ([]*UsersPool, error) {
	var resp []*UsersPool
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `is_deleted` = '0'", usersPoolRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	return resp, err
}

func (m *customUsersPoolModel) FindByUidAndAid(ctx context.Context, userId, activityId int64) ([]*UsersPool, error) {
	var resp []*UsersPool
	query := fmt.Sprintf("select %s from %s `activity_id` = ? and `user_id` = ?  and `is_deleted` = '0'", usersPoolRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	return resp, err
}
