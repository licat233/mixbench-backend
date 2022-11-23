package ctxdata

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	CtxKeyJwtUserId     = "jwtUserId"
	CtxKeyJwtActivityId = "jwtActivityId"
	CtxKeyJwtIp         = "jwtIp"
	CtxKeyJwtPlatform   = "jwtPlatform"
)

// 获取用户id
func GetUidFromCtx(ctx context.Context) int64 {
	var user_id int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			user_id = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err: %+v", err)
		}
	}
	return user_id
}

// 获取活动id
func GetAidFromCtx(ctx context.Context) int64 {
	var activity_id int64
	if jsonAid, ok := ctx.Value(CtxKeyJwtActivityId).(json.Number); ok {
		if int64Aid, err := jsonAid.Int64(); err == nil {
			activity_id = int64Aid
		} else {
			logx.WithContext(ctx).Errorf("GetAidFromCtx err: %+v", err)
		}
	}
	return activity_id
}
