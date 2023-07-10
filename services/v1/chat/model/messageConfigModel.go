package model

import (
	"context"
	"fmt"

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
		FindsByMsgkeyProjectId(ctx context.Context, msgkey string, projectId int64) ([]*MessageConfig, error)
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

func (m *customMessageConfigModel) FindsByMsgkeyProjectId(ctx context.Context, msgkey string, projectId int64) ([]*MessageConfig, error) {
	var resp = make([]*MessageConfig, 0)
	query := fmt.Sprintf("select %s from %s where `msgkey` = ? and `project_id` = ? ", messageConfigRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, msgkey, projectId)
	return resp, err
}
