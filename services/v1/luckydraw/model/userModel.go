package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel  // extended interface by gozero
		user_model // extended interface by genzero

		FindsByName(ctx context.Context, name string) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	//用于插入数据的字段
	userRowsExpectAutoSet = strings.Join(stringx.Remove(userFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	cacheUserIdPrefix = "cache:luckydraw:user:id:" // modifying cache id prefix by genzero
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *customUserModel) FindsByName(ctx context.Context, name string) ([]*User, error) {
	var resp []*User
	query := fmt.Sprintf("select %s from %s where `name` like '%%%s%%' ", userRows, m.table, name)
	fmt.Println(query)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
