package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IpblackModel = (*customIpblackModel)(nil)

type (
	// IpblackModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIpblackModel.
	IpblackModel interface {
		ipblackModel  // extended interface by gozero
		ipblack_model // extended interface by genzero
	}

	customIpblackModel struct {
		*defaultIpblackModel
	}
)

// NewIpblackModel returns a model for the database table.
func NewIpblackModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IpblackModel {
	return &customIpblackModel{
		defaultIpblackModel: newIpblackModel(conn, c, opts...),
	}
}
