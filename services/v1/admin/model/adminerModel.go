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
		CountByEmail(ctx context.Context, email string) (int64, error)
	}

	customAdminerModel struct {
		*defaultAdminerModel
	}
)

// NewAdminerModel returns a model for the database table.
func NewAdminerModel(conn sqlx.SqlConn, c cache.CacheConf) AdminerModel {
	cacheAdminerIdPrefix = "cache:admin:adminer:id:" // modifying cache id prefix by genzero
	return &customAdminerModel{
		defaultAdminerModel: newAdminerModel(conn, c),
	}
}

func (m *customAdminerModel) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	query := fmt.Sprintf("select count(*) as count from %s where `email` = ? limit 1", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, email)
	return count != 0, err
}

func (m *customAdminerModel) CountByEmail(ctx context.Context, email string) (int64, error) {
	var res int64
	query := fmt.Sprintf("select count(*) from %s where `email` = ?", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &res, query, email)
	return res, err
}
