package model

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ PrizeModel = (*customPrizeModel)(nil)
var PrizeTableName = "prize"
var cacheLuckDrawPrizeUuidPrefix = "cache:luckDraw:prize:uuid:"

type (
	// PrizeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeModel.
	PrizeModel interface {
		prizeModel
		DeleteSoft(ctx context.Context, data *Prize) error
		FindOneByUuid(ctx context.Context, uuid string) (*Prize, error)

		FindCount(ctx context.Context) (int64, error)
		FindPrizesByActivityUuid(ctx context.Context, activityUuid string) ([]*Prize, error)
		FindsByUserAndActivity(ctx context.Context, userId, activityId int64) ([]*Prize, error)
		FindsByUserAndActivityAndStatus(ctx context.Context, userId, activityId, Status int64) ([]*Prize, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string, prize *Prize) (resp []*Prize, total int64, err error)
		FindInIds(ctx context.Context, prizesIds []int64) ([]*Prize, error)
		TableName() string
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

func (m *customPrizeModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawPrizeUuidPrefix, uuid)
}

// 软删除
func (m *customPrizeModel) DeleteSoft(ctx context.Context, data *Prize) error {
	luckDrawPrizeIdKey := m.formatPrimary(data.Id)
	luckDrawPrizeUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawPrizeIdKey, luckDrawPrizeUuidKey)
	//完成
	return err
}

func (m *customPrizeModel) FindOneByUuid(ctx context.Context, uuid string) (*Prize, error) {
	luckDrawPrizeUuidKey := m.formatUuidKey(uuid)
	var resp Prize
	err := m.QueryRowCtx(ctx, &resp, luckDrawPrizeUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", prizeRows, m.table)
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

func (m *customPrizeModel) FindOne(ctx context.Context, id int64) (*Prize, error) {
	luckDrawPrizeIdKey := fmt.Sprintf("%s%v", cacheLuckDrawPrizeIdPrefix, id)
	var resp Prize
	err := m.QueryRowCtx(ctx, &resp, luckDrawPrizeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ?  and `is_deleted` = '0' limit 1", prizeRows, m.table)
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

func (m *customPrizeModel) Insert(ctx context.Context, data *Prize) (sql.Result, error) {
	luckDrawPrizeIdKey := m.formatPrimary(data.Id)
	luckDrawPrizeUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, prizeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Image, data.Price, data.Status)
	}, luckDrawPrizeIdKey, luckDrawPrizeUuidKey)
	return ret, err
}

func (m *customPrizeModel) Update(ctx context.Context, data *Prize) error {
	luckDrawPrizeIdKey := m.formatPrimary(data.Id)
	luckDrawPrizeUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?  and `is_deleted` = '0'", m.table, prizeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Image, data.Price, data.Status, data.Id)
	}, luckDrawPrizeIdKey, luckDrawPrizeUuidKey)
	return err
}

//----------------------------------------------------------------

func (m *customPrizeModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `is_deleted`= '0'", m.table)
	var total int64
	err := m.QueryRowNoCacheCtx(ctx, &total, query)
	switch err {
	case nil:
		return total, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *customPrizeModel) FindPrizesByActivityUuid(ctx context.Context, activityUuid string) ([]*Prize, error) {
	// newPrizeRows := resetTableRows(prizeFieldNames, "p")
	query := fmt.Sprintf("select p.* from %s as p left join %s as pp on p.id = pp.prize_id left join %s as a on pp.activity_id = a.id where p.is_deleted = 0 and pp.is_deleted = 0 and a.is_deleted = 0 and a.uuid = ?", PrizeTableName, PrizesPoolTableName, ActivityTableName)
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
	//select * from prize as p left join prizes_pool as pp ON pp.prize_id = p.id LEFT JOIN lucker as l ON l.prizes_pool_id = pp.id WHERE pp.activity_id = 1 AND p.is_deleted = 0 AND pp.is_deleted = 0 AND l.is_deleted = 0;
	// newPrizeRows := resetTableRows(prizeFieldNames, "p")
	query := fmt.Sprintf("select p.* from %s as p left join %s as pp on pp.prize_id = p.id left join %s as l on l.prizes_pool_id = pp.id where l.user_id = ? and pp.activity_id = ? AND p.is_deleted = 0 AND pp.is_deleted = 0 AND l.is_deleted = 0", m.table, PrizesPoolTableName, PrizeTableName)
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
	//select * from prize as p left join prizes_pool as pp ON pp.prize_id = p.id LEFT JOIN lucker as l ON l.prizes_pool_id = pp.id WHERE pp.activity_id = 1 AND p.is_deleted = 0 AND pp.is_deleted = 0 AND l.is_deleted = 0;
	// newPrizeRows := resetTableRows(prizeFieldNames, "p")
	query := fmt.Sprintf("select p.* from %s as p left join %s as pp on pp.prize_id = p.id left join %s as l on l.prizes_pool_id = pp.id where l.user_id = ? and pp.activity_id = ? and l.status = ? AND p.is_deleted = 0 AND pp.is_deleted = 0 AND l.is_deleted = 0", m.table, PrizesPoolTableName, PrizeTableName)
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

func (m *customPrizeModel) FindList(ctx context.Context, pageSize, page int64, keyword string, prize *Prize) (resp []*Prize, total int64, err error) {
	if pageSize == 0 {
		return
	}
	sq := squirrel.Select(prizeRows).From(m.table).Where("`is_deleted`= '0'")
	sqCount := squirrel.Select("COUNT(*)").From(m.table).Where("`is_deleted`= '0'")
	var mark bool
	if prize != nil {
		//name的优先级大于keyword
		if prize.Name != "" {
			sq = sq.Where("name LIKE ?", fmt.Sprint("%", prize.Name, "%"))
			sqCount = sqCount.Where("name LIKE ?", fmt.Sprint("%", prize.Name, "%"))
			mark = true
		}
		if prize.Price != -1 {
			sq = sq.Where("`price` = ?", prize.Price)
			sqCount = sqCount.Where("`price` = ?", prize.Price)
			mark = true
		}
		if prize.Status != -1 {
			sq = sq.Where("`status` = ?", prize.Status)
			sqCount = sqCount.Where("`status` = ?", prize.Status)
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

func (m *customPrizeModel) FindInIds(ctx context.Context, prizesIds []int64) ([]*Prize, error) {
	if len(prizesIds) == 0 {
		return nil, nil
	}
	var temps []string
	for _, id := range prizesIds {
		temps = append(temps, strconv.Itoa(int(id)))
	}
	whereSql := strings.Join(temps, ",")
	query := fmt.Sprintf("select %s from %s where `is_deleted` = '0' and `id` in(%s)", prizeRows, m.table, whereSql)
	var resp []*Prize
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

func (m *customPrizeModel) TableName() string {
	return m.table
}
