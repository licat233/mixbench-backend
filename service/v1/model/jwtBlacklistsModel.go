package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ JwtBlacklistsModel = (*customJwtBlacklistsModel)(nil)
var JwtBlacklistsTableName = "jwt_blacklists"
var cacheLuckDrawJwtBlacklistsUuidPrefix = "cache:luckDraw:jwtBlacklists:uuid:"
var cacheLuckDrawJwtBlacklistsTokenPrefix = "cache:luckDraw:jwtBlacklists:token:"

type (
	// JwtBlacklistsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJwtBlacklistsModel.
	JwtBlacklistsModel interface {
		jwtBlacklistsModel
		DeleteOne(ctx context.Context, id int64) error
		FindOneByUuid(ctx context.Context, uuid string) (*JwtBlacklists, error)
		FindOneByToken(ctx context.Context, token string) (*JwtBlacklists, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64) ([]*JwtBlacklists, error)
		Exist(ctx context.Context, id int64) (bool, error)
		ExistByToken(ctx context.Context, token string) (bool, error)
	}

	customJwtBlacklistsModel struct {
		*defaultJwtBlacklistsModel
	}
)

// NewJwtBlacklistsModel returns a model for the database table.
func NewJwtBlacklistsModel(conn sqlx.SqlConn, c cache.CacheConf) JwtBlacklistsModel {
	return &customJwtBlacklistsModel{
		defaultJwtBlacklistsModel: newJwtBlacklistsModel(conn, c),
	}
}

func (m *customJwtBlacklistsModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawJwtBlacklistsUuidPrefix, uuid)
}

func (m *customJwtBlacklistsModel) formatTokenKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawJwtBlacklistsTokenPrefix, uuid)
}

func (m *customJwtBlacklistsModel) DeleteOne(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil && err != ErrNotFound {
		return err
	}
	if err == ErrNotFound {
		return nil
	}
	luckDrawJwtBlacklistsIdKey := m.formatPrimary(data.Id)
	luckDrawJwtBlacklistsUuidKey := m.formatTokenKey(data.Uuid)
	luckDrawJwtBlacklistsTokenKey := m.formatTokenKey(data.Token)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, luckDrawJwtBlacklistsIdKey, luckDrawJwtBlacklistsUuidKey, luckDrawJwtBlacklistsTokenKey)
	return err
}

func (m *customJwtBlacklistsModel) FindOneByUuid(ctx context.Context, uuid string) (*JwtBlacklists, error) {
	luckDrawJwtBlacklistsUuidKey := m.formatUuidKey(uuid)
	var resp JwtBlacklists
	err := m.QueryRowCtx(ctx, &resp, luckDrawJwtBlacklistsUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", jwtBlacklistsRows, m.table)
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

func (m *customJwtBlacklistsModel) FindOneByToken(ctx context.Context, token string) (*JwtBlacklists, error) {
	luckDrawJwtBlacklistsTokenKey := m.formatTokenKey(token)
	var resp JwtBlacklists
	err := m.QueryRowCtx(ctx, &resp, luckDrawJwtBlacklistsTokenKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `token` = ? limit 1", jwtBlacklistsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, token)
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

func (m *customJwtBlacklistsModel) Insert(ctx context.Context, data *JwtBlacklists) (sql.Result, error) {
	luckDrawJwtBlacklistsIdKey := m.formatPrimary(data.Id)
	luckDrawJwtBlacklistsUuidKey := m.formatTokenKey(data.Uuid)
	luckDrawJwtBlacklistsTokenKey := m.formatTokenKey(data.Token)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, jwtBlacklistsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Token, data.Platform, data.Ip, data.ExpireAt, data.CreateAt, data.UpdateAt)
	}, luckDrawJwtBlacklistsUuidKey, luckDrawJwtBlacklistsIdKey, luckDrawJwtBlacklistsTokenKey)
	return ret, err
}

func (m *customJwtBlacklistsModel) Update(ctx context.Context, data *JwtBlacklists) error {
	luckDrawJwtBlacklistsIdKey := m.formatPrimary(data.Id)
	luckDrawJwtBlacklistsUuidKey := m.formatTokenKey(data.Uuid)
	luckDrawJwtBlacklistsTokenKey := m.formatTokenKey(data.Token)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, jwtBlacklistsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Token, data.Platform, data.Ip, data.ExpireAt, data.CreateAt, data.UpdateAt, data.Id)
	}, luckDrawJwtBlacklistsUuidKey, luckDrawJwtBlacklistsIdKey, luckDrawJwtBlacklistsTokenKey)
	return err
}

//--

func (m *customJwtBlacklistsModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	return count, err
}

func (m *customJwtBlacklistsModel) FindList(ctx context.Context, pageSize, page int64) ([]*JwtBlacklists, error) {
	if pageSize == 0 {
		return []*JwtBlacklists{}, nil
	}
	query := fmt.Sprintf("select %s from %s limit %v,%v", jwtBlacklistsRows, m.table, (page-1)*pageSize, pageSize)
	if pageSize < 0 || page < 0 {
		query = fmt.Sprintf("select %s from %s where `is_deleted`= '0'", projectRows, m.table)
	}
	var resp []*JwtBlacklists
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

func (m *customJwtBlacklistsModel) Exist(ctx context.Context, id int64) (bool, error) {
	_, err := m.FindOne(ctx, id)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	return err != ErrNotFound, nil
}

func (m *customJwtBlacklistsModel) ExistByToken(ctx context.Context, token string) (bool, error) {
	_, err := m.FindOneByToken(ctx, token)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	return err != ErrNotFound, nil
}
