package model

import (
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	_ ActivityModel = (*customActivityModel)(nil)
)

type (
	// ActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActivityModel.
	ActivityModel interface {
		activityModel  // extended interface by gozero
		activity_model // extended interface by genzero
		PatternEnumList() []*Enum
	}

	customActivityModel struct {
		*defaultActivityModel
	}
)

// NewActivityModel returns a model for the database table.
func NewActivityModel(conn sqlx.SqlConn, c cache.CacheConf) ActivityModel {
	//用于插入数据的字段
	activityRowsExpectAutoSet = strings.Join(stringx.Remove(activityFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	activityRowsWithPlaceHolder = strings.Join(stringx.Remove(activityFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customActivityModel{
		defaultActivityModel: newActivityModel(conn, c),
	}
}

func (m *customActivityModel) PatternEnumList() []*Enum {
	return []*Enum{
		{
			Label: "会员制",
			Value: 0,
		},
		{
			Label: "公开制",
			Value: 1,
		},
		{
			Label: "VIP制",
			Value: 2,
		},
	}
}
