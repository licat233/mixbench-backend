package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ JwtBlacklistModel = (*customJwtBlacklistModel)(nil)
var cachejwtBlacklistTokenPrefix = "cache:jwtBlacklist:token:"

type (
	// JwtBlacklistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJwtBlacklistModel.
	JwtBlacklistModel interface {
		jwtBlacklistModel  // extended interface by gozero
		jwtBlacklist_model // extended interface by genzero
		Insert(ctx context.Context, data *JwtBlacklist) (sql.Result, error)
		Update(ctx context.Context, data *JwtBlacklist) error
		Delete(ctx context.Context, id int64) error
		FindCount(ctx context.Context) (int64, error)
		Exist(ctx context.Context, id int64) (bool, error)
		FindOneByToken(ctx context.Context, token string) (*JwtBlacklist, error)
		ExistByToken(ctx context.Context, token string) (bool, error)
	}

	customJwtBlacklistModel struct {
		*defaultJwtBlacklistModel
	}
)

// NewJwtBlacklistModel returns a model for the database table.
func NewJwtBlacklistModel(conn sqlx.SqlConn, c cache.CacheConf) JwtBlacklistModel {
	return &customJwtBlacklistModel{
		defaultJwtBlacklistModel: newJwtBlacklistModel(conn, c),
	}
}

func (m *customJwtBlacklistModel) Insert(ctx context.Context, data *JwtBlacklist) (sql.Result, error) {
	jwtBlacklistIdKey := fmt.Sprintf("%s%v", cacheJwtBlacklistIdPrefix, data.Id)
	jwtBlacklistTokenKey := m.formatTokenKey(data.Token)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, jwtBlacklistRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Token, data.Platform, data.Ip, data.ExpireAt)
	}, jwtBlacklistIdKey, jwtBlacklistTokenKey)
	return ret, err
}

func (m *customJwtBlacklistModel) Update(ctx context.Context, data *JwtBlacklist) error {
	jwtBlacklistIdKey := fmt.Sprintf("%s%v", cacheJwtBlacklistIdPrefix, data.Id)
	jwtBlacklistTokenKey := m.formatTokenKey(data.Token)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, jwtBlacklistRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Token, data.Platform, data.Ip, data.ExpireAt, data.Id)
	}, jwtBlacklistIdKey, jwtBlacklistTokenKey)
	return err
}

func (m *customJwtBlacklistModel) Delete(ctx context.Context, id int64) error {
	jwtBlacklistIdKey := fmt.Sprintf("%s%v", cacheJwtBlacklistIdPrefix, id)
	data, err := m.FindOne(ctx, id)
	if err != nil && err != ErrNotFound {
		return err
	}
	if err == ErrNotFound {
		return nil
	}
	jwtBlacklistTokenKey := m.formatTokenKey(data.Token)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, jwtBlacklistIdKey, jwtBlacklistTokenKey)
	return err
}

func (m *customJwtBlacklistModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	return count, err
}

func (m *customJwtBlacklistModel) Exist(ctx context.Context, id int64) (bool, error) {
	_, err := m.FindOne(ctx, id)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	return err != ErrNotFound, nil
}

func (m *customJwtBlacklistModel) formatTokenKey(uuid string) string {
	return fmt.Sprintf("%s%v", cachejwtBlacklistTokenPrefix, uuid)
}

func (m *customJwtBlacklistModel) FindOneByToken(ctx context.Context, token string) (*JwtBlacklist, error) {
	jwtBlacklistTokenKey := m.formatTokenKey(token)
	var resp JwtBlacklist
	err := m.QueryRowCtx(ctx, &resp, jwtBlacklistTokenKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `token` = ? limit 1", jwtBlacklistRows, m.table)
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

func (m *customJwtBlacklistModel) ExistByToken(ctx context.Context, token string) (bool, error) {
	_, err := m.FindOneByToken(ctx, token)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	return err != ErrNotFound, nil
}
