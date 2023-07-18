package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TokenBlacklistModel = (*customTokenBlacklistModel)(nil)
var cacheTokenBlacklistTokenPrefix = "cache:tokenBlacklist:token:"

type (
	// TokenBlacklistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTokenBlacklistModel.
	TokenBlacklistModel interface {
		tokenBlacklistModel  // extended interface by gozero
		tokenBlacklist_model // extended interface by genzero
		Insert(ctx context.Context, data *TokenBlacklist) (sql.Result, error)
		Update(ctx context.Context, data *TokenBlacklist) error
		Delete(ctx context.Context, id int64) error
		FindCount(ctx context.Context) (int64, error)
		Exist(ctx context.Context, id int64) (bool, error)
		FindOneByToken(ctx context.Context, token string) (*TokenBlacklist, error)
		ExistsByToken(ctx context.Context, token string) (bool, error)
	}

	customTokenBlacklistModel struct {
		*defaultTokenBlacklistModel
	}
)

// NewTokenBlacklistModel returns a model for the database table.
func NewTokenBlacklistModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TokenBlacklistModel {
	return &customTokenBlacklistModel{
		defaultTokenBlacklistModel: newTokenBlacklistModel(conn, c, opts...),
	}
}

func (m *customTokenBlacklistModel) Insert(ctx context.Context, data *TokenBlacklist) (sql.Result, error) {
	if data.Token = strings.TrimSpace(data.Token); data.Token == "" {
		return nil, ErrNilToken
	}
	tokenBlacklistIdKey := m.formatPrimary(data.Id)
	tokenBlacklistTokenKey := m.formatTokenKey(data.Token)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, tokenBlacklistRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Token, data.Platform, data.Ip)
	}, tokenBlacklistIdKey, tokenBlacklistTokenKey)
	return ret, err
}

func (m *customTokenBlacklistModel) Update(ctx context.Context, data *TokenBlacklist) error {
	if data.Token = strings.TrimSpace(data.Token); data.Token == "" {
		return ErrNilToken
	}
	tokenBlacklistIdKey := m.formatPrimary(data.Id)
	tokenBlacklistTokenKey := m.formatTokenKey(data.Token)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tokenBlacklistRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Token, data.Platform, data.Ip, data.Id)
	}, tokenBlacklistIdKey, tokenBlacklistTokenKey)
	return err
}

func (m *customTokenBlacklistModel) Delete(ctx context.Context, id int64) error {
	tokenBlacklistIdKey := fmt.Sprintf("%s%v", cacheTokenBlacklistIdPrefix, id)
	data, err := m.FindById(ctx, id)
	if err != nil && err != ErrNotFound {
		return err
	}
	tokenBlacklistTokenKey := ""
	if err == nil {
		tokenBlacklistTokenKey = m.formatTokenKey(data.Token)
	}
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tokenBlacklistIdKey, tokenBlacklistTokenKey)
	return err
}

func (m *customTokenBlacklistModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	return count, err
}

func (m *customTokenBlacklistModel) Exist(ctx context.Context, id int64) (bool, error) {
	_, err := m.FindById(ctx, id)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	return err != ErrNotFound, nil
}

func (m *customTokenBlacklistModel) formatTokenKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheTokenBlacklistTokenPrefix, uuid)
}

func (m *customTokenBlacklistModel) ExistsByToken(ctx context.Context, token string) (bool, error) {
	tokenBlacklistTokenKey := m.formatTokenKey(token)
	var count int64
	err := m.QueryRowCtx(ctx, &count, tokenBlacklistTokenKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select count(*) as count from %s where `token` = ? limit 1", m.table)
		return conn.QueryRowCtx(ctx, v, query, token)
	})
	return count > 0, err
}

func (m *customTokenBlacklistModel) FindOneByToken(ctx context.Context, token string) (*TokenBlacklist, error) {
	tokenBlacklistTokenKey := m.formatTokenKey(token)
	var resp TokenBlacklist
	err := m.QueryRowCtx(ctx, &resp, tokenBlacklistTokenKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `token` = ? limit 1", tokenBlacklistRows, m.table)
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
