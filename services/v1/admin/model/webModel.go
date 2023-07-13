package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WebModel = (*customWebModel)(nil)

type (
	// WebModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWebModel.
	WebModel interface {
		webModel  // extended interface by gozero
		web_model // extended interface by genzero
	}

	customWebModel struct {
		*defaultWebModel
	}
)

// NewWebModel returns a model for the database table.
func NewWebModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WebModel {
	return &customWebModel{
		defaultWebModel: newWebModel(conn, c, opts...),
	}
}
