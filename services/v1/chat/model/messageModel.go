package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel  // extended interface by gozero
		message_model // extended interface by genzero

		Insert(ctx context.Context, data *Message) (sql.Result, error)
		DeleteExpiredRecords(ctx context.Context) error
		FindsToVisitorByVisitorIds(ctx context.Context, visitorIds []int64) ([]*Message, error)
		FindsToKefuByKefuIds(ctx context.Context, kefuIds []int64) ([]*Message, error)
		FindsToKefuByKefuId(ctx context.Context, kefuId int64) ([]*Message, error)
		FindsToVisitorByVisitorId(ctx context.Context, visitorId int64) ([]*Message, error)
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn, c, opts...),
	}
}

func (m *customMessageModel) Insert(ctx context.Context, data *Message) (sql.Result, error) {
	ret, err := m.defaultMessageModel.Insert(ctx, data)
	if err != nil {
		return ret, err
	}
	// 设置缓存过期时间
	messageIdKey := fmt.Sprintf("%s%v", cacheMessageIdPrefix, data.Id)
	jsonData, _ := json.Marshal(*data)
	exp := 3 * 24 * time.Hour
	err = m.SetCacheWithExpireCtx(ctx, messageIdKey, jsonData, exp)
	return ret, err
}

func (m *customMessageModel) DeleteExpiredRecords(ctx context.Context) error {
	query := fmt.Sprintf("delete from %s where expiration_time <= NOW()", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query)
	return err
}

func (m *customMessageModel) FindsToKefuByKefuIds(ctx context.Context, kefuIds []int64) ([]*Message, error) {
	var resp = make([]*Message, 0)
	if len(kefuIds) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `kefu_id` in(?) and `msg_type` = '0' and `is_deleted` = '0' ", messageRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, kefuIds)
	return resp, err
}

func (m *customMessageModel) FindsToVisitorByVisitorIds(ctx context.Context, visitorIds []int64) ([]*Message, error) {
	var resp = make([]*Message, 0)
	if len(visitorIds) == 0 {
		return resp, nil
	}
	query := fmt.Sprintf("select %s from %s where `visitor_id` in(?) and `msg_type` = '1' and `is_deleted` = '0' ", messageRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, visitorIds)
	return resp, err
}

// 查找游客发给客服的消息
func (m *customMessageModel) FindsToKefuByKefuId(ctx context.Context, kefuId int64) ([]*Message, error) {
	var resp = make([]*Message, 0)
	query := fmt.Sprintf("select %s from %s where `kefu_id` = ? and `msg_type` = '0' and `is_deleted` = '0' ", messageRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, kefuId)
	return resp, err
}

// 查找客服发给游客的信息
func (m *customMessageModel) FindsToVisitorByVisitorId(ctx context.Context, visitorId int64) ([]*Message, error) {
	var resp = make([]*Message, 0)
	query := fmt.Sprintf("select %s from %s where `visitor_id` = ? and `msg_type` = '1' and `is_deleted` = '0' ", messageRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, visitorId)
	return resp, err
}
