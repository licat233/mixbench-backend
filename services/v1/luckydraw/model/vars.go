package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

var (
	//插入时omit的字段
	BaseInsertOmitField = []string{"`id`", "`create_at`", "`update_at`", "`is_deleted`", "`delete_at`"}
	//更新时omit的字段
	BaseUpdateOmitField = []string{"`id`", "`uuid`", "`create_at`", "`update_at`", "`is_deleted`", "`delete_at`"}
)

type ListReq struct {
	PageSize int64
	Page     int64
	Keyword  string
}

type Enum struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}
