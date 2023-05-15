package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ UserPoolModel = (*customUserPoolModel)(nil)

type (
	// UserPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPoolModel.
	UserPoolModel interface {
		userPoolModel  // extended interface by gozero
		userPool_model // extended interface by genzero

		FindByUidAndAid(ctx context.Context, userId, activityId int64) ([]*UserPool, error)
	}

	customUserPoolModel struct {
		*defaultUserPoolModel
	}
)

// NewUserPoolModel returns a model for the database table.
func NewUserPoolModel(conn sqlx.SqlConn, c cache.CacheConf) UserPoolModel {
	//用于插入数据的字段
	userPoolRowsExpectAutoSet = strings.Join(stringx.Remove(userPoolFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	userPoolRowsWithPlaceHolder = strings.Join(stringx.Remove(userPoolFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customUserPoolModel{
		defaultUserPoolModel: newUserPoolModel(conn, c),
	}
}

func (m *customUserPoolModel) FindByUidAndAid(ctx context.Context, userId, activityId int64) ([]*UserPool, error) {
	var resp []*UserPool
	query := fmt.Sprintf("select %s from %s `activity_id` = ? and `user_id` = ?  and `is_deleted` = '0'", userPoolRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	return resp, err
}
