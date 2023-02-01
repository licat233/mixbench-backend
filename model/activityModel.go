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
	_                       ActivityModel = (*customActivityModel)(nil)
	ActivityTableName                     = "activity"
	cacheActivityUuidPrefix               = "cache:activity:uuid:"
)

type (
	// ActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActivityModel.
	ActivityModel interface {
		activityModel
		DeleteSoft(ctx context.Context, data *Activity) error
		FindOneByUuid(ctx context.Context, uuid string) (*Activity, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, activity *Activity) (resp []*Activity, total int64, err error)
		FindsByProjectId(ctx context.Context, projectId int64) ([]*Activity, error)
		TableName() string
		FindsByIds(ctx context.Context, ids []int64) ([]*Activity, error)
		FindsByPattern(ctx context.Context, pattern int64) ([]*Activity, error)
		PatternEnumList() []*Enum
	}

	customActivityModel struct {
		*defaultActivityModel
	}
)

// NewActivityModel returns a model for the database table.
func NewActivityModel(conn sqlx.SqlConn, c cache.CacheConf) ActivityModel {
	//用于插入数据的字段
	activityRowsExpectAutoSet = strings.Join(stringx.Remove(activityFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	activityRowsWithPlaceHolder = strings.Join(stringx.Remove(activityFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customActivityModel{
		defaultActivityModel: newActivityModel(conn, c),
	}
}

func (m *customActivityModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheActivityUuidPrefix, uuid)
}

// 软删除
func (m *customActivityModel) DeleteSoft(ctx context.Context, data *Activity) error {
	activityIdKey := m.formatPrimary(data.Id)
	activityUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, activityIdKey, activityUuidKey)
	//完成
	return err
}

func (m *customActivityModel) FindOneByUuid(ctx context.Context, uuid string) (*Activity, error) {
	activityUuidKey := m.formatUuidKey(uuid)
	var resp Activity
	err := m.QueryRowCtx(ctx, &resp, activityUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", activityRows, m.table)
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

func (m *customActivityModel) Insert(ctx context.Context, data *Activity) (sql.Result, error) {
	activityIdKey := m.formatPrimary(data.Id)
	activityUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, activityRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.ProjectId, data.Status, data.Pattern, data.MaxCount, data.MaxPrizes, data.StartTime, data.EndTime)
	}, activityIdKey, activityUuidKey)
	return ret, err
}

func (m *customActivityModel) Update(ctx context.Context, data *Activity) error {
	activityIdKey := m.formatPrimary(data.Id)
	activityUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, activityRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.ProjectId, data.Status, data.Pattern, data.MaxCount, data.MaxPrizes, data.StartTime, data.EndTime, data.Id)
	}, activityIdKey, activityUuidKey)
	return err
}

//----------

func (m *customActivityModel) FindCount(ctx context.Context) (int64, error) {
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

func (m *customActivityModel) FindList(ctx context.Context, pageSize, page int64, keyword string, activity *Activity) (resp []*Activity, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(activityRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if activity != nil {
		//name的优先级大于keyword
		if activity.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", activity.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", activity.Name, "%"))
			mark = true
		}
		if utils.IdValid(activity.ProjectId) {
			sq = sq.Where("`project_id` = ?", activity.ProjectId)
			sqCount = sqCount.Where("`project_id` = ?", activity.ProjectId)
			mark = true
		}
		if activity.Status != -1 {
			sq = sq.Where("`status` = ?", activity.Status)
			sqCount = sqCount.Where("`status` = ?", activity.Status)
			mark = true
		}
		if activity.Pattern != -1 {
			sq = sq.Where("`pattern` = ?", activity.Pattern)
			sqCount = sqCount.Where("`pattern` = ?", activity.Pattern)
			mark = true
		}
		if activity.MaxCount != -2 {
			sq = sq.Where("`max_count` = ?", activity.MaxCount)
			sqCount = sqCount.Where("`max_count` = ?", activity.MaxCount)
			mark = true
		}
		if activity.MaxPrizes != -2 {
			sq = sq.Where("`max_prizes` = ?", activity.MaxPrizes)
			sqCount = sqCount.Where("`max_prizes` = ?", activity.MaxPrizes)
			mark = true
		}
		if !activity.StartTime.IsZero() && !activity.EndTime.IsZero() {
			sq = sq.Where(squirrel.And{
				squirrel.GtOrEq{
					"start_time": activity.StartTime,
				},
				squirrel.LtOrEq{
					"end_time": activity.EndTime,
				},
			})
			sqCount = sqCount.Where(squirrel.And{
				squirrel.GtOrEq{
					"start_time": activity.StartTime,
				},
				squirrel.LtOrEq{
					"end_time": activity.EndTime,
				},
			})
			mark = true
		}
		if !activity.StartTime.IsZero() && activity.EndTime.IsZero() {
			sq = sq.Where(squirrel.GtOrEq{
				"start_time": activity.StartTime,
			})
			sqCount = sqCount.Where(squirrel.GtOrEq{
				"start_time": activity.StartTime,
			})
			mark = true
		}
		if activity.StartTime.IsZero() && !activity.EndTime.IsZero() {
			sq = sq.Where(squirrel.LtOrEq{
				"end_time": activity.EndTime,
			})
			sqCount = sqCount.Where(squirrel.LtOrEq{
				"end_time": activity.EndTime,
			})
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

func (m *customActivityModel) FindsByProjectId(ctx context.Context, projectId int64) ([]*Activity, error) {
	var resp []*Activity
	var err error

	if projectId == -1 {
		query := fmt.Sprintf("select %s from %s where `is_deleted`= '0'", activityRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query)
	} else {
		query := fmt.Sprintf("select %s from %s where `is_deleted`= '0' and `project_id` = ?", activityRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, projectId)
	}

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customActivityModel) TableName() string {
	return m.table
}

func (m *customActivityModel) FindsByIds(ctx context.Context, ids []int64) ([]*Activity, error) {
	var resp []*Activity
	where := utils.SliceIntToString(ids, ",")
	query := fmt.Sprintf("select %s from %s where `id` in(%s) and `is_deleted` = '0'", activityRows, m.table, where)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	return resp, err
}

func (m *customActivityModel) FindsByPattern(ctx context.Context, pattern int64) ([]*Activity, error) {
	var resp []*Activity
	query := fmt.Sprintf("select %s from %s where `pattern` = ? and `is_deleted` = '0'", activityRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, pattern)
	return resp, err
}

func (m *customActivityModel) PatternEnumList() []*Enum {
	return []*Enum{
		{
			Label: "会员制",
			Value: 0,
		},
		{
			Label: "公开制",
			Value: 1,
		},
		{
			Label: "VIP制",
			Value: 2,
		},
	}
}
