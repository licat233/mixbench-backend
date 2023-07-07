package wshub

import (
	"fmt"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/types"
)

func NewPublicMessage(content string, kefuId, visitorId int64) *types.Message {
	return &types.Message{
		Id:         0,
		Uuid:       uniqueid.NewUUID(),
		Content:    content,
		VisitorId:  visitorId,
		KefuId:     kefuId,
		MsgType:    PublicMsgType,
		Status:     0,
		IsRead:     0,
		IsSend:     0,
		SendTime:   0,
		ExpireTime: 0,
	}
}

func NewSystemMessage(content any, kefuId, visitorId int64) *types.Message {
	return &types.Message{
		Id:         0,
		Uuid:       uniqueid.NewUUID(),
		Content:    fmt.Sprintf("%s", content),
		VisitorId:  visitorId,
		KefuId:     kefuId,
		MsgType:    SystemMsgType,
		Status:     0,
		IsRead:     0,
		IsSend:     0,
		SendTime:   0,
		ExpireTime: 0,
	}
}
