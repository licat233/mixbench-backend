package model

import (
	"context"
	"fmt"
	"mixbench/common/utils"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ PrizePoolModel = (*customPrizePoolModel)(nil)

type (
	// PrizePoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizePoolModel.
	PrizePoolModel interface {
		prizePoolModel  // extended interface by gozero
		prizePool_model // extended interface by genzero

		FindByActivityUuid(ctx context.Context, activityUuid string) ([]*PrizePool, error)
		FindsByActivityIds(ctx context.Context, activityIds []int64) ([]*PrizePool, error)

		FindsByActivityAndPrize(ctx context.Context, activityId int64, prize_id int64) ([]*PrizePool, error)
		Conn() sqlc.CachedConn
	}

	customPrizePoolModel struct {
		*defaultPrizePoolModel
	}
)

// NewPrizePoolModel returns a model for the database table.
func NewPrizePoolModel(conn sqlx.SqlConn, c cache.CacheConf) PrizePoolModel {
	//用于插入数据的字段
	prizePoolRowsExpectAutoSet = strings.Join(stringx.Remove(prizePoolFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	prizePoolRowsWithPlaceHolder = strings.Join(stringx.Remove(prizePoolFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	cachePrizePoolIdPrefix = "cache:luckydraw:prizePool:id:" // modifying cache id prefix by genzero
	return &customPrizePoolModel{
		defaultPrizePoolModel: newPrizePoolModel(conn, c),
	}
}

// 获取某活动奖品池内的奖品
// TODO：为了避免复杂性，一般先依据activityUUID查询到activity_id，再根据activity_id进行查询
// TODO：但在此处，我们直接使用关联查询，一劳永逸
func (m *customPrizePoolModel) FindByActivityUuid(ctx context.Context, activityUuid string) ([]*PrizePool, error) {
	query := fmt.Sprintf("SELECT pp.* FROM %s AS pp LEFT JOIN %s AS a ON pp.activity_id = a.id WHERE a.uuid = ? ;", m.table, ActivityTableName)
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

func (m *customPrizePoolModel) FindsByActivityIds(ctx context.Context, activityIds []int64) ([]*PrizePool, error) {
	var resp []*PrizePool
	where := utils.SliceIntToString(activityIds, ",")
	query := fmt.Sprintf("select %s from %s where `activity_id` in(%s) ", prizePoolRows, m.table, where)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	return resp, err
}

func (m *customPrizePoolModel) FindsByActivityAndPrize(ctx context.Context, activityId, prizeId int64) ([]*PrizePool, error) {
	var resp []*PrizePool
	query := fmt.Sprintf("select %s from %s where `activity_id` = ? and `prize_id` = ? ", prizePoolRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, activityId, prizeId)
	return resp, err
}

func (m *customPrizePoolModel) Conn() sqlc.CachedConn {
	return m.CachedConn
}
