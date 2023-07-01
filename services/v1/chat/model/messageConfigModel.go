package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageConfigModel = (*customMessageConfigModel)(nil)

type (
	// MessageConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageConfigModel.
	MessageConfigModel interface {
		messageConfigModel  // extended interface by gozero
		messageConfig_model // extended interface by genzero
		UpdateContent(ctx context.Context, Id int64, content string) error
	}

	customMessageConfigModel struct {
		*defaultMessageConfigModel
	}
)

// NewMessageConfigModel returns a model for the database table.
func NewMessageConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MessageConfigModel {
	return &customMessageConfigModel{
		defaultMessageConfigModel: newMessageConfigModel(conn, c, opts...),
	}
}

func (m *customMessageConfigModel) UpdateContent(ctx context.Context, Id int64, content string) error {
	data, err := m.FindOne(ctx, Id)
	if err != nil {
		return err
	}
	data.Content = content
	return m.Update(ctx, data)
}
