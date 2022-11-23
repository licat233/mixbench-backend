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

var _ UserModel = (*customUserModel)(nil)
var UserTableName = "user"
var cacheLuckDrawUserUuidPrefix = "cache:luckDraw:user:uuid:"

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		DeleteSoft(ctx context.Context, data *User) error
		FindOneByUuid(ctx context.Context, uuid string) (*User, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, user *User) (resp []*User, total int64, err error)
		Exists(ctx context.Context, id int64) (bool, *User, error)
		FindByPasspord(ctx context.Context, passport string) (*User, error)
		FindByPhone(ctx context.Context, phone string) (*User, error)
		FindByEmail(ctx context.Context, email string) (*User, error)
		FindsByName(ctx context.Context, name string) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	//用于插入数据的字段
	userRowsExpectAutoSet = strings.Join(stringx.Remove(userFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *customUserModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawUserUuidPrefix, uuid)
}

// 软删除
func (m *customUserModel) DeleteSoft(ctx context.Context, data *User) error {
	luckDrawUserIdKey := m.formatPrimary(data.Id)
	luckDrawUserUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawUserIdKey, luckDrawUserUuidKey)
	//完成
	return err
}

func (m *customUserModel) FindOneByUuid(ctx context.Context, uuid string) (*User, error) {
	luckDrawUserUuidKey := m.formatUuidKey(uuid)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, luckDrawUserUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", userRows, m.table)
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

func (m *customUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	luckDrawUserIdKey := m.formatPrimary(id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, luckDrawUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' limit 1", userRows, m.table)
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

func (m *customUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	luckDrawUserIdKey := m.formatPrimary(data.Id)
	luckDrawUserUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Passport, data.Email, data.Phone, data.Name, data.Address, data.Score, data.Status)
	}, luckDrawUserIdKey, luckDrawUserUuidKey)
	return ret, err
}

func (m *customUserModel) Update(ctx context.Context, data *User) error {
	luckDrawUserIdKey := m.formatPrimary(data.Id)
	luckDrawUserUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0' ", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Passport, data.Email, data.Phone, data.Name, data.Address, data.Score, data.Status, data.Id)
	}, luckDrawUserIdKey, luckDrawUserUuidKey)
	return err
}

//----------------------------------------------------------------

func (m *customUserModel) FindCount(ctx context.Context) (int64, error) {
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

func (m *customUserModel) FindList(ctx context.Context, pageSize, page int64, keyword string, user *User) (resp []*User, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(userRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if user != nil {
		//name的优先级大于keyword
		if user.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", user.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", user.Name, "%"))
			mark = true
		}
		if user.Passport != "" {
			sq = sq.Where("`passport` = ?", user.Passport)
			sqCount = sqCount.Where("`passport` = ?", user.Passport)
			mark = true
		}
		if user.Email != "" {
			sq = sq.Where("`email` = ?", user.Email)
			sqCount = sqCount.Where("`email` = ?", user.Email)
			mark = true
		}
		if user.Phone != "" {
			sq = sq.Where("`phone` = ?", user.Phone)
			sqCount = sqCount.Where("`phone` = ?", user.Phone)
			mark = true
		}
		if user.Address != "" {
			sq = sq.Where("`address` = ?", user.Address)
			sqCount = sqCount.Where("`address` = ?", user.Address)
			mark = true
		}
		if user.Score != -1 {
			sq = sq.Where("`score` = ?", user.Score)
			sqCount = sqCount.Where("`score` = ?", user.Score)
			mark = true
		}
		if user.Status != -1 {
			sq = sq.Where("`status` = ?", user.Status)
			sqCount = sqCount.Where("`status` = ?", user.Status)
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

func (m *customUserModel) Exists(ctx context.Context, id int64) (bool, *User, error) {
	user, err := m.FindOne(ctx, id)
	if err != nil && err != ErrNotFound {
		return false, nil, err
	}
	if err == ErrNotFound {
		return false, nil, nil
	}
	// if user == nil {
	// 	return false, nil, errors.New("has found, but user == nil")
	// }
	if user.IsDeleted != 1 {
		return false, nil, nil
	}

	return true, user, nil
}

func (m *customUserModel) FindByPasspord(ctx context.Context, passport string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `passport` = ? and `is_deleted` = '0' limit 1", userRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, passport)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) FindByEmail(ctx context.Context, email string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `email` = ? and `is_deleted` = '0' limit 1", userRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) FindByPhone(ctx context.Context, phone string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `phone` = ? and `is_deleted` = '0' limit 1", userRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) FindsByName(ctx context.Context, name string) ([]*User, error) {
	var resp []*User
	query := fmt.Sprintf("select %s from %s where `name` like '%%%s%%' and `is_deleted` = '0'", userRows, m.table, name)
	fmt.Println(query)
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
