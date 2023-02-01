package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsTagModel = (*customGoodsTagModel)(nil)

type (
	// GoodsTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsTagModel.
	GoodsTagModel interface {
		goodsTagModel
	}

	customGoodsTagModel struct {
		*defaultGoodsTagModel
	}
)

// NewGoodsTagModel returns a model for the database table.
func NewGoodsTagModel(conn sqlx.SqlConn, c cache.CacheConf) GoodsTagModel {
	return &customGoodsTagModel{
		defaultGoodsTagModel: newGoodsTagModel(conn, c),
	}
}
