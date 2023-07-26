package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProjectModel = (*customProjectModel)(nil)

type (
	// ProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectModel.
	ProjectModel interface {
		projectModel  // extended interface by gozero
		project_model // extended interface by genzero
	}

	customProjectModel struct {
		*defaultProjectModel
	}
)

// NewProjectModel returns a model for the database table.
func NewProjectModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProjectModel {
	cacheProjectIdPrefix = "cache:admin:project:id:" // modifying cache id prefix by genzero
	return &customProjectModel{
		defaultProjectModel: newProjectModel(conn, c, opts...),
	}
}
