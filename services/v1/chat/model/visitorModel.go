package model

import (
	"context"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VisitorModel = (*customVisitorModel)(nil)

type (
	// VisitorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVisitorModel.
	VisitorModel interface {
		visitorModel  // extended interface by gozero
		visitor_model // extended interface by genzero

		FindsByKefuid(ctx context.Context, kefuId int64, pageSize, page int64) (resp []*Visitor, total int64, err error)
	}

	customVisitorModel struct {
		*defaultVisitorModel
	}
)

// NewVisitorModel returns a model for the database table.
func NewVisitorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VisitorModel {
	return &customVisitorModel{
		defaultVisitorModel: newVisitorModel(conn, c, opts...),
	}
}

func (m *customVisitorModel) FindsByKefuid(ctx context.Context, kefuId int64, pageSize, page int64) (resp []*Visitor, total int64, err error) {
	if kefuId <= 0 {
		return
	}
	sq := squirrel.Select(visitorRows).From(m.table).Where("`is_deleted`= '0'").Where("`kefu_id` = ?", kefuId).OrderBy("status")

	if pageSize > 0 && page > 0 {
		sqCount := sq.RemoveLimit().RemoveOffset()
		sq = sq.Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize))
		queryCount, agrsCount, e := sqCount.ToSql()
		if e != nil {
			err = e
			return
		}
		queryCount = strings.ReplaceAll(queryCount, visitorRows, "COUNT(*)")
		if err = m.QueryRowNoCacheCtx(ctx, &total, queryCount, agrsCount...); err != nil {
			return
		}
	}
	query, agrs, err := sq.ToSql()
	if err != nil {
		return
	}
	resp = make([]*Visitor, 0)
	if err = m.QueryRowsNoCacheCtx(ctx, &resp, query, agrs...); err != nil {
		return
	}
	return
}
