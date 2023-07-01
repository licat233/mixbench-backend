package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ KefuModel = (*customKefuModel)(nil)

type (
	// KefuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customKefuModel.
	KefuModel interface {
		kefuModel  // extended interface by gozero
		kefu_model // extended interface by genzero
	}

	customKefuModel struct {
		*defaultKefuModel
	}
)

// NewKefuModel returns a model for the database table.
func NewKefuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) KefuModel {
	return &customKefuModel{
		defaultKefuModel: newKefuModel(conn, c, opts...),
	}
}
