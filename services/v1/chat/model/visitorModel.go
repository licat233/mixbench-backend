package model

import (
	"context"

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
	cacheVisitorIdPrefix = "cache:chat:visitor:id:" // modifying cache id prefix by genzero
	return &customVisitorModel{
		defaultVisitorModel: newVisitorModel(conn, c, opts...),
	}
}

func (m *customVisitorModel) FindsByKefuid(ctx context.Context, kefuId int64, pageSize, page int64) (resp []*Visitor, total int64, err error) {
	if kefuId <= 0 {
		return
	}
	sq := squirrel.Select(visitorRows).From(m.table).Where("`is_deleted`= '0'").Where("`kefu_id` = ?", kefuId).OrderBy("status").OrderBy("last_login_at")

	if pageSize > 0 && page > 0 {
		sq = sq.Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize))
		sqCount := sq.RemoveLimit().RemoveOffset()
		queryCount, agrsCount, e := sqCount.RemoveColumns().Column("COUNT(*)").ToSql()
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
	resp = make([]*Visitor, 0)
	if err = m.QueryRowsNoCacheCtx(ctx, &resp, query, agrs...); err != nil {
		return
	}
	return
}
