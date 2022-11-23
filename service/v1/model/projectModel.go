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

var _ ProjectModel = (*customProjectModel)(nil)
var ProjectTableName = "project"
var cacheLuckDrawProjectUuidPrefix = "cache:luckDraw:project:uuid:"

type (
	// ProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectModel.
	ProjectModel interface {
		projectModel
		DeleteSoft(ctx context.Context, data *Project) error
		FindOneByUuid(ctx context.Context, uuid string) (*Project, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, project *Project) (resp []*Project, total int64, err error)
		TableName() string
	}

	customProjectModel struct {
		*defaultProjectModel
	}
)

// NewProjectModel returns a model for the database table.
func NewProjectModel(conn sqlx.SqlConn, c cache.CacheConf) ProjectModel {
	//用于插入数据的字段
	projectRowsExpectAutoSet = strings.Join(stringx.Remove(projectFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	projectRowsWithPlaceHolder = strings.Join(stringx.Remove(projectFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customProjectModel{
		defaultProjectModel: newProjectModel(conn, c),
	}
}

func (m *customProjectModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawProjectUuidPrefix, uuid)
}

// 软删除
func (m *customProjectModel) DeleteSoft(ctx context.Context, data *Project) error {
	luckDrawProjectIdKey := m.formatPrimary(data.Id)
	luckDrawProjectUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawProjectIdKey, luckDrawProjectUuidKey)
	//完成
	return err
}

func (m *customProjectModel) FindOneByUuid(ctx context.Context, uuid string) (*Project, error) {
	luckDrawProjectUuidKey := m.formatUuidKey(uuid)
	var resp Project
	err := m.QueryRowCtx(ctx, &resp, luckDrawProjectUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", projectRows, m.table)
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

func (m *customProjectModel) FindOne(ctx context.Context, id int64) (*Project, error) {
	luckDrawProjectIdKey := m.formatPrimary(id)
	var resp Project
	err := m.QueryRowCtx(ctx, &resp, luckDrawProjectIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' limit 1", projectRows, m.table)
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

func (m *customProjectModel) Insert(ctx context.Context, data *Project) (sql.Result, error) {
	luckDrawProjectIdKey := m.formatPrimary(data.Id)
	luckDrawProjectUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, projectRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Status)
	}, luckDrawProjectIdKey, luckDrawProjectUuidKey)
	return ret, err
}

func (m *customProjectModel) Update(ctx context.Context, data *Project) error {
	luckDrawProjectIdKey := m.formatPrimary(data.Id)
	luckDrawProjectUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, projectRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Status, data.Id)
	}, luckDrawProjectIdKey, luckDrawProjectUuidKey)
	return err
}

//-------------------

func (m *customProjectModel) FindCount(ctx context.Context) (int64, error) {
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

func (m *customProjectModel) FindList(ctx context.Context, pageSize, page int64, keyword string, project *Project) (resp []*Project, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(projectRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if project != nil {
		//name的优先级大于keyword
		if project.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", project.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", project.Name, "%"))
			mark = true
		}
		if project.Status != -1 {
			sq = sq.Where("`status` = ?", project.Status)
			sqCount = sqCount.Where("`status` = ?", project.Status)
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

func (m *customProjectModel) TableName() string {
	return m.table
}
