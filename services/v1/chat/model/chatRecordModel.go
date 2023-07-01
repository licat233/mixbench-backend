package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatRecordModel = (*customChatRecordModel)(nil)

type (
	// ChatRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatRecordModel.
	ChatRecordModel interface {
		chatRecordModel  // extended interface by gozero
		chatRecord_model // extended interface by genzero
	}

	customChatRecordModel struct {
		*defaultChatRecordModel
	}
)

// NewChatRecordModel returns a model for the database table.
func NewChatRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ChatRecordModel {
	return &customChatRecordModel{
		defaultChatRecordModel: newChatRecordModel(conn, c, opts...),
	}
}
