package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminerModel = (*customAdminerModel)(nil)

type (
	// AdminerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminerModel.
	AdminerModel interface {
		adminerModel  // extended interface by gozero
		adminer_model // extended interface by genzero
		ExistsByEmail(ctx context.Context, email string) (bool, error)
	}

	customAdminerModel struct {
		*defaultAdminerModel
	}
)

// NewAdminerModel returns a model for the database table.
func NewAdminerModel(conn sqlx.SqlConn, c cache.CacheConf) AdminerModel {
	return &customAdminerModel{
		defaultAdminerModel: newAdminerModel(conn, c),
	}
}

func (m *customAdminerModel) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	query := fmt.Sprintf("select count(*) as count from %s where `email` = ? limit 1", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &count, query, email)
	return count != 0, err
}
