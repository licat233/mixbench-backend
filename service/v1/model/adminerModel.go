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

var _ AdminerModel = (*customAdminerModel)(nil)
var AdminerTableName = "adminer"
var cacheLuckDrawAdminerUuidPrefix = "cache:luckDraw:adminer:uuid:"

type (
	// AdminerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminerModel.
	AdminerModel interface {
		adminerModel
		DeleteSoft(ctx context.Context, data *Adminer) error
		FindOneByUuid(ctx context.Context, uuid string) (*Adminer, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, adminer *Adminer) ([]*Adminer, int64, error)
		FindOneByEmail(ctx context.Context, email string) (*Adminer, error)
		FindOneByPassport(ctx context.Context, passport string) (*Adminer, error)

		ExsitByEmail(ctx context.Context, email string) (bool, error)
	}

	customAdminerModel struct {
		*defaultAdminerModel
	}
)

// NewAdminerModel returns a model for the database table.
func NewAdminerModel(conn sqlx.SqlConn, c cache.CacheConf) AdminerModel {
	//用于插入数据的字段
	adminerRowsExpectAutoSet = strings.Join(stringx.Remove(adminerFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	adminerRowsWithPlaceHolder = strings.Join(stringx.Remove(adminerFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customAdminerModel{
		defaultAdminerModel: newAdminerModel(conn, c),
	}
}

func (m *customAdminerModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawAdminerUuidPrefix, uuid)
}

// 软删除
func (m *customAdminerModel) DeleteSoft(ctx context.Context, data *Adminer) error {
	luckDrawAdminerIdKey := m.formatPrimary(data.Id)
	luckDrawAdminerUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawAdminerIdKey, luckDrawAdminerUuidKey)
	//完成
	return err
}

func (m *customAdminerModel) FindOneByUuid(ctx context.Context, uuid string) (*Adminer, error) {
	luckDrawAdminerUuidKey := m.formatUuidKey(uuid)
	var resp Adminer
	err := m.QueryRowCtx(ctx, &resp, luckDrawAdminerUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
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

func (m *customAdminerModel) Insert(ctx context.Context, data *Adminer) (sql.Result, error) {
	luckDrawAdminerIdKey := m.formatPrimary(data.Id)
	luckDrawAdminerUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, adminerRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Avatar, data.Passport, data.Password, data.Email, data.Status, data.IsSuperAdmin, data.LoginCount, data.LastLogin)
	}, luckDrawAdminerIdKey, luckDrawAdminerUuidKey)
	return ret, err
}

func (m *customAdminerModel) Update(ctx context.Context, data *Adminer) error {
	luckDrawAdminerIdKey := m.formatPrimary(data.Id)
	luckDrawAdminerUuidKey := m.formatUuidKey(data.Uuid)
	query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, adminerRowsWithPlaceHolder)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query, data.Name, data.Avatar, data.Passport, data.Password, data.Email, data.Status, data.IsSuperAdmin, data.LoginCount, data.LastLogin, data.Id)
	}, luckDrawAdminerIdKey, luckDrawAdminerUuidKey)
	return err
}

//-------

func (m *customAdminerModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `is_deleted`= '0'", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	//就算不存在，err也不会等于 model.ErrNotFound
	return count, err
}

func (m *customAdminerModel) FindList(ctx context.Context, pageSize, page int64, keyword string, adminer *Adminer) (resp []*Adminer, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(adminerRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if adminer != nil {
		//name的优先级大于keyword
		if adminer.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", adminer.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", adminer.Name, "%"))
			mark = true
		}
		if adminer.Passport != "" {
			sq = sq.Where("`passport` = ?", adminer.Passport)
			sqCount = sqCount.Where("`passport` = ?", adminer.Passport)
			mark = true
		}
		if adminer.Email != "" {
			sq = sq.Where("`email` = ?", adminer.Email)
			sqCount = sqCount.Where("`email` = ?", adminer.Email)
			mark = true
		}
		// if adminer.Avatar != "" {
		// 	sq.Where("`avatar` = ?", adminer.Avatar)
		// 	mark = true
		// }
		if adminer.Password != "" {
			sq = sq.Where("`password` = ?", adminer.Password)
			sqCount = sqCount.Where("`password` = ?", adminer.Password)
			mark = true
		}
		if adminer.Status != -1 {
			sq = sq.Where("`status` = ?", adminer.Status)
			sqCount = sqCount.Where("`status` = ?", adminer.Status)
			mark = true
		}
		if adminer.IsSuperAdmin != -1 {
			sq = sq.Where("`is_super_admin` = ?", adminer.IsSuperAdmin)
			sqCount = sqCount.Where("`is_super_admin` = ?", adminer.IsSuperAdmin)
			mark = true
		}
		if adminer.LoginCount != -1 {
			sq = sq.Where("`login_count` = ?", adminer.LoginCount)
			sqCount = sqCount.Where("`login_count` = ?", adminer.LoginCount)
			mark = true
		}
		if !adminer.LastLogin.IsZero() {
			dayStr := adminer.LastLogin.Format("2006-01-02")
			sq = sq.Where("DATE_FORMAT(`last_login`,'%Y-%m-%d') = ?", dayStr)
			sqCount = sqCount.Where("DATE_FORMAT(`last_login`,'%Y-%m-%d') = ?", dayStr)
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

func (m *customAdminerModel) FindOneByPassport(ctx context.Context, passport string) (*Adminer, error) {
	query := fmt.Sprintf("select %s from %s where `passport` = ? and `is_deleted`= '0' limit 1", adminerRows, m.table)
	var resp Adminer
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

func (m *customAdminerModel) FindOneByEmail(ctx context.Context, email string) (*Adminer, error) {
	query := fmt.Sprintf("select %s from %s where `email` = ? and `is_deleted`= '0' limit 1", adminerRows, m.table)
	var resp Adminer
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

func (m *customAdminerModel) ExsitByEmail(ctx context.Context, email string) (bool, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `email` = ? and `is_deleted`= 0 limit 1", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query, email)
	//就算不存在，err也不会等于 model.ErrNotFound
	// fmt.Printf("count: %+v , error: %+v \n", count, err)
	return count != 0, err
}
