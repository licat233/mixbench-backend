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

var _ jwtBlacklistModel = (*customjwtBlacklistModel)(nil)
var jwtBlacklistTableName = "jwt_blacklist"
var cachejwtBlacklistUuidPrefix = "cache:jwtBlacklist:uuid:"
var cachejwtBlacklistTokenPrefix = "cache:jwtBlacklist:token:"

type (
	// jwtBlacklistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customjwtBlacklistModel.
	JwtBlacklistModel interface {
		jwtBlacklistModel
		DeleteOne(ctx context.Context, id int64) error
		FindOneByUuid(ctx context.Context, uuid string) (*JwtBlacklist, error)
		FindOneByToken(ctx context.Context, token string) (*JwtBlacklist, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64) ([]*JwtBlacklist, error)
		Exist(ctx context.Context, id int64) (bool, error)
		ExistByToken(ctx context.Context, token string) (bool, error)
	}

	customjwtBlacklistModel struct {
		*defaultjwtBlacklistModel
	}
)

// NewJwtBlacklistModel returns a model for the database table.
func NewJwtBlacklistModel(conn sqlx.SqlConn, c cache.CacheConf) JwtBlacklistModel {
	//用于插入数据的字段
	jwtBlacklistRowsExpectAutoSet = strings.Join(stringx.Remove(jwtBlacklistFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	jwtBlacklistRowsWithPlaceHolder = strings.Join(stringx.Remove(jwtBlacklistFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customjwtBlacklistModel{
		defaultjwtBlacklistModel: newjwtBlacklistModel(conn, c),
	}
}

func (m *customjwtBlacklistModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cachejwtBlacklistUuidPrefix, uuid)
}

func (m *customjwtBlacklistModel) formatTokenKey(uuid string) string {
	return fmt.Sprintf("%s%v", cachejwtBlacklistTokenPrefix, uuid)
}

func (m *customjwtBlacklistModel) DeleteOne(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil && err != ErrNotFound {
		return err
	}
	if err == ErrNotFound {
		return nil
	}
	jwtBlacklistIdKey := m.formatPrimary(data.Id)
	jwtBlacklistUuidKey := m.formatUuidKey(data.Uuid)
	jwtBlacklistTokenKey := m.formatTokenKey(data.Token)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, jwtBlacklistIdKey, jwtBlacklistUuidKey, jwtBlacklistTokenKey)
	return err
}

func (m *customjwtBlacklistModel) FindOneByUuid(ctx context.Context, uuid string) (*JwtBlacklist, error) {
	jwtBlacklistUuidKey := m.formatUuidKey(uuid)
	var resp JwtBlacklist
	err := m.QueryRowCtx(ctx, &resp, jwtBlacklistUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", jwtBlacklistRows, m.table)
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

func (m *customjwtBlacklistModel) FindOneByToken(ctx context.Context, token string) (*JwtBlacklist, error) {
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

func (m *customjwtBlacklistModel) Insert(ctx context.Context, data *JwtBlacklist) (sql.Result, error) {
	jwtBlacklistIdKey := m.formatPrimary(data.Id)
	jwtBlacklistUuidKey := m.formatUuidKey(data.Uuid)
	jwtBlacklistTokenKey := m.formatTokenKey(data.Token)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, jwtBlacklistRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Token, data.Platform, data.Ip, data.ExpireAt)
	}, jwtBlacklistUuidKey, jwtBlacklistIdKey, jwtBlacklistTokenKey)
	return ret, err
}

func (m *customjwtBlacklistModel) Update(ctx context.Context, data *JwtBlacklist) error {
	jwtBlacklistIdKey := m.formatPrimary(data.Id)
	jwtBlacklistUuidKey := m.formatUuidKey(data.Uuid)
	jwtBlacklistTokenKey := m.formatTokenKey(data.Token)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, jwtBlacklistRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Token, data.Platform, data.Ip, data.ExpireAt, data.Id)
	}, jwtBlacklistUuidKey, jwtBlacklistIdKey, jwtBlacklistTokenKey)
	return err
}

//---------------------------------------------------------------------------------

func (m *customjwtBlacklistModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	return count, err
}

func (m *customjwtBlacklistModel) FindList(ctx context.Context, pageSize, page int64) ([]*JwtBlacklist, error) {
	if pageSize == 0 {
		return []*JwtBlacklist{}, nil
	}
	query := fmt.Sprintf("select %s from %s limit %v,%v", jwtBlacklistRows, m.table, (page-1)*pageSize, pageSize)
	if pageSize < 0 || page < 0 {
		query = fmt.Sprintf("select %s from %s where `is_deleted`= '0'", projectRows, m.table)
	}
	var resp []*JwtBlacklist
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

func (m *customjwtBlacklistModel) Exist(ctx context.Context, id int64) (bool, error) {
	_, err := m.FindOne(ctx, id)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	return err != ErrNotFound, nil
}

func (m *customjwtBlacklistModel) ExistByToken(ctx context.Context, token string) (bool, error) {
	_, err := m.FindOneByToken(ctx, token)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	return err != ErrNotFound, nil
}
