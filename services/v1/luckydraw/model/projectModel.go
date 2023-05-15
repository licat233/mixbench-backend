package model

import (
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
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
func NewProjectModel(conn sqlx.SqlConn, c cache.CacheConf) ProjectModel {
	//用于插入数据的字段
	projectRowsExpectAutoSet = strings.Join(stringx.Remove(projectFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	projectRowsWithPlaceHolder = strings.Join(stringx.Remove(projectFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customProjectModel{
		defaultProjectModel: newProjectModel(conn, c),
	}
}
