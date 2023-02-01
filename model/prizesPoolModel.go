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

var _ PrizePoolModel = (*customPrizePoolModel)(nil)
var PrizePoolTableName = "prize_pool"
var cachePrizePoolUuidPrefix = "cache:prizesPool:uuid:"

type (
	// PrizePoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizePoolModel.
	PrizePoolModel interface {
		prizesPoolModel
		QueryOne(ctx context.Context, id int64) (*PrizePool, error)
		DeleteSoft(ctx context.Context, data *PrizePool) error
		FindOneByUuid(ctx context.Context, uuid string) (*PrizePool, error)

		FindCount(ctx context.Context) (int64, error)
		FindByActivityUuid(ctx context.Context, activityUuid string) ([]*PrizePool, error)
		FindList(ctx context.Context, pageSize, page int64, prizesPool *PrizePool) (resp []*PrizePool, total int64, err error)
		FindsByActivityId(ctx context.Context, activityId int64) ([]*PrizePool, error)
		FindsByActivityIds(ctx context.Context, activityIds []int64) ([]*PrizePool, error)

		FindsByActivityAndPrize(ctx context.Context, activityId int64, prize_id int64) ([]*PrizePool, error)
		TableName() string
		Conn() sqlc.CachedConn
	}

	customPrizePoolModel struct {
		*defaultPrizePoolModel
	}
)

// NewPrizePoolModel returns a model for the database table.
func NewPrizePoolModel(conn sqlx.SqlConn, c cache.CacheConf) PrizePoolModel {
	//用于插入数据的字段
	prizesPoolRowsExpectAutoSet = strings.Join(stringx.Remove(prizesPoolFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	prizesPoolRowsWithPlaceHolder = strings.Join(stringx.Remove(prizesPoolFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customPrizePoolModel{
		defaultPrizePoolModel: newPrizePoolModel(conn, c),
	}
}

func (m *customPrizePoolModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cachePrizePoolUuidPrefix, uuid)
}

// 软删除
func (m *customPrizePoolModel) DeleteSoft(ctx context.Context, data *PrizePool) error {
	prizesPoolIdKey := m.formatPrimary(data.Id)
	prizesPoolUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, prizesPoolIdKey, prizesPoolUuidKey)
	//完成
	return err
}

func (m *customPrizePoolModel) FindOneByUuid(ctx context.Context, uuid string) (*PrizePool, error) {
	prizesPoolUuidKey := m.formatUuidKey(uuid)
	var resp PrizePool
	err := m.QueryRowCtx(ctx, &resp, prizesPoolUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", prizesPoolRows, m.table)
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

func (m *customPrizePoolModel) FindOne(ctx context.Context, id int64) (*PrizePool, error) {
	prizesPoolIdKey := m.formatPrimary(id)
	var resp PrizePool
	err := m.QueryRowCtx(ctx, &resp, prizesPoolIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' limit 1", prizesPoolRows, m.table)
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

func (m *customPrizePoolModel) QueryOne(ctx context.Context, id int64) (*PrizePool, error) {
	prizesPoolIdKey := m.formatPrimary(id)
	var resp PrizePool
	err := m.QueryRowCtx(ctx, &resp, prizesPoolIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", prizesPoolRows, m.table)
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

func (m *customPrizePoolModel) Insert(ctx context.Context, data *PrizePool) (sql.Result, error) {
	prizesPoolIdKey := m.formatPrimary(data.Id)
	prizesPoolUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, prizesPoolRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.ActivityId, data.PrizeId, data.GradeId, data.LuckValue, data.Status, data.Total, data.Remainder)
	}, prizesPoolIdKey, prizesPoolUuidKey)
	return ret, err
}

func (m *customPrizePoolModel) Update(ctx context.Context, data *PrizePool) error {
	prizesPoolIdKey := m.formatPrimary(data.Id)
	prizesPoolUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, prizesPoolRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ActivityId, data.PrizeId, data.GradeId, data.LuckValue, data.Status, data.Total, data.Remainder, data.Id)
	}, prizesPoolIdKey, prizesPoolUuidKey)
	return err
}

//----------------------------------------------------------------

func (m *customPrizePoolModel) TableName() string {
	return m.table
}

func (m *customPrizePoolModel) FindCount(ctx context.Context) (int64, error) {
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

func (m *customPrizePoolModel) FindsByActivityId(ctx context.Context, activityId int64) ([]*PrizePool, error) {
	var resp []*PrizePool
	var err error

	if activityId == -1 {
		query := fmt.Sprintf("select %s from %s where `is_deleted` = '0'", prizesPoolRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query)
	} else {
		query := fmt.Sprintf("select %s from %s where `activity_id` = ? and `is_deleted` = '0'", prizesPoolRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, activityId)
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

// 获取某活动奖品池内的奖品
// TODO：为了避免复杂性，一般先依据activityUUID查询到activity_id，再根据activity_id进行查询
// TODO：但在此处，我们直接使用关联查询，一劳永逸
func (m *customPrizePoolModel) FindByActivityUuid(ctx context.Context, activityUuid string) ([]*PrizePool, error) {
	query := fmt.Sprintf("SELECT pp.* FROM %s AS pp LEFT JOIN %s AS a ON pp.activity_id = a.id WHERE a.uuid = ? AND pp.is_deleted = 0 AND a.is_deleted = 0;", m.table, ActivityTableName)
	var resp []*PrizePool
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, activityUuid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPrizePoolModel) FindList(ctx context.Context, pageSize, page int64, prizesPool *PrizePool) (resp []*PrizePool, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sqQuery := squirrel.Select(prizesPoolRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	if prizesPool != nil {
		if prizesPool.ActivityId != -1 {
			sqQuery = sqQuery.Where("`activity_id` = ?", prizesPool.ActivityId)
			sqCount = sqCount.Where("`activity_id` = ?", prizesPool.ActivityId)
		}
		if prizesPool.PrizeId != -1 {
			sqQuery = sqQuery.Where("`prize_id` = ?", prizesPool.PrizeId)
			sqCount = sqCount.Where("`prize_id` = ?", prizesPool.PrizeId)
		}
		if prizesPool.GradeId != -1 {
			sqQuery = sqQuery.Where("`grade_id` = ?", prizesPool.GradeId)
			sqCount = sqCount.Where("`grade_id` = ?", prizesPool.GradeId)
		}
		if prizesPool.LuckValue != -1 {
			sqQuery = sqQuery.Where("`luck_value` = ?", prizesPool.LuckValue)
			sqCount = sqCount.Where("`luck_value` = ?", prizesPool.LuckValue)
		}
		if prizesPool.Status != -1 {
			sqQuery = sqQuery.Where("`status` = ?", prizesPool.Status)
			sqCount = sqCount.Where("`status` = ?", prizesPool.Status)
		}
		if prizesPool.Total != -1 {
			sqQuery = sqQuery.Where("`number` = ?", prizesPool.Total)
			sqCount = sqCount.Where("`number` = ?", prizesPool.Total)
		}
		if prizesPool.Remainder != -1 {
			sqQuery = sqQuery.Where("`remainder` = ?", prizesPool.Remainder)
			sqCount = sqCount.Where("`remainder` = ?", prizesPool.Remainder)
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

func (m *customPrizePoolModel) FindsByActivityIds(ctx context.Context, activityIds []int64) ([]*PrizePool, error) {
	var resp []*PrizePool
	where := utils.SliceIntToString(activityIds, ",")
	query := fmt.Sprintf("select %s from %s where `activity_id` in(%s) and `is_deleted` = '0'", prizesPoolRows, m.table, where)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	return resp, err
}

func (m *customPrizePoolModel) FindsByActivityAndPrize(ctx context.Context, activityId, prizeId int64) ([]*PrizePool, error) {
	var resp []*PrizePool
	query := fmt.Sprintf("select %s from %s where `activity_id` = ? and `prize_id` = ? and `is_deleted` = '0'", prizesPoolRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, activityId, prizeId)
	return resp, err
}

func (m *customPrizePoolModel) Conn() sqlc.CachedConn {
	return m.CachedConn
}
