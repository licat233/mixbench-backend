package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IpBlacklistModel = (*customIpBlacklistModel)(nil)

type (
	// IpBlacklistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIpBlacklistModel.
	IpBlacklistModel interface {
		ipBlacklistModel  // extended interface by gozero
		ipBlacklist_model // extended interface by genzero
	}

	customIpBlacklistModel struct {
		*defaultIpBlacklistModel
	}
)

// NewIpBlacklistModel returns a model for the database table.
func NewIpBlacklistModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IpBlacklistModel {
	return &customIpBlacklistModel{
		defaultIpBlacklistModel: newIpBlacklistModel(conn, c, opts...),
	}
}
