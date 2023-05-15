package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ PrizeModel = (*customPrizeModel)(nil)

type (
	// PrizeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeModel.
	PrizeModel interface {
		prizeModel  // extended interface by gozero
		prize_model // extended interface by genzero

		FindPrizesByActivityUuid(ctx context.Context, activityUuid string) ([]*Prize, error)
		FindsByUserAndActivity(ctx context.Context, userId, activityId int64) ([]*Prize, error)
		FindsByUserAndActivityAndStatus(ctx context.Context, userId, activityId, Status int64) ([]*Prize, error)
	}

	customPrizeModel struct {
		*defaultPrizeModel
	}
)

// NewPrizeModel returns a model for the database table.
func NewPrizeModel(conn sqlx.SqlConn, c cache.CacheConf) PrizeModel {
	//用于插入数据的字段
	prizeRowsExpectAutoSet = strings.Join(stringx.Remove(prizeFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	prizeRowsWithPlaceHolder = strings.Join(stringx.Remove(prizeFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customPrizeModel{
		defaultPrizeModel: newPrizeModel(conn, c),
	}
}

func (m *customPrizeModel) FindPrizesByActivityUuid(ctx context.Context, activityUuid string) ([]*Prize, error) {
	// newPrizeRows := resetTableRows(prizeFieldNames, "p")
	query := fmt.Sprintf("select p.* from %s as p left join %s as pp on p.id = pp.prize_id left join %s as a on pp.activity_id = a.id where p.is_deleted = 0 and pp.is_deleted = 0 and a.is_deleted = 0 and a.uuid = ?", PrizeTableName, PrizePoolTableName, ActivityTableName)
	var resp []*Prize
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

// FindsByUserAndActivity 查找某用户在某活动中的奖品
func (m *customPrizeModel) FindsByUserAndActivity(ctx context.Context, userId, activityId int64) ([]*Prize, error) {
	//select * from prize as p left join prize_pool as pp ON pp.prize_id = p.id LEFT JOIN lucky as l ON l.prize_pool_id = pp.id WHERE pp.activity_id = 1 AND p.is_deleted = 0 AND pp.is_deleted = 0 AND l.is_deleted = 0;
	// newPrizeRows := resetTableRows(prizeFieldNames, "p")
	query := fmt.Sprintf("select p.* from %s as p left join %s as pp on pp.prize_id = p.id left join %s as l on l.prize_pool_id = pp.id where l.user_id = ? and pp.activity_id = ? AND p.is_deleted = 0 AND pp.is_deleted = 0 AND l.is_deleted = 0", m.table, PrizePoolTableName, PrizeTableName)
	var resp []*Prize
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId, activityId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindsByUserAndActivityAndStatus 查找某用户在某活动中的奖品
func (m *customPrizeModel) FindsByUserAndActivityAndStatus(ctx context.Context, userId, activityId, Status int64) ([]*Prize, error) {
	//select * from prize as p left join prize_pool as pp ON pp.prize_id = p.id LEFT JOIN lucky as l ON l.prize_pool_id = pp.id WHERE pp.activity_id = 1 AND p.is_deleted = 0 AND pp.is_deleted = 0 AND l.is_deleted = 0;
	// newPrizeRows := resetTableRows(prizeFieldNames, "p")
	query := fmt.Sprintf("select p.* from %s as p left join %s as pp on pp.prize_id = p.id left join %s as l on l.prize_pool_id = pp.id where l.user_id = ? and pp.activity_id = ? and l.status = ? AND p.is_deleted = 0 AND pp.is_deleted = 0 AND l.is_deleted = 0", m.table, PrizePoolTableName, PrizeTableName)
	var resp []*Prize
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId, activityId, Status)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
