package model

import (
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ PrizeGradeModel = (*customPrizeGradeModel)(nil)

type (
	// PrizeGradeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeGradeModel.
	PrizeGradeModel interface {
		prizeGradeModel  // extended interface by gozero
		prizeGrade_model // extended interface by genzero

	}

	customPrizeGradeModel struct {
		*defaultPrizeGradeModel
	}
)

// NewPrizeGradeModel returns a model for the database table.
func NewPrizeGradeModel(conn sqlx.SqlConn, c cache.CacheConf) PrizeGradeModel {
	//用于插入数据的字段
	prizeGradeRowsExpectAutoSet = strings.Join(stringx.Remove(prizeGradeFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	prizeGradeRowsWithPlaceHolder = strings.Join(stringx.Remove(prizeGradeFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customPrizeGradeModel{
		defaultPrizeGradeModel: newPrizeGradeModel(conn, c),
	}
}
