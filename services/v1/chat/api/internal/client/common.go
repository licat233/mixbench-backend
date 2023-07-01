package client

import (
	"fmt"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/types"
)

func NewPublicMessage(content string) *types.Message {
	return &types.Message{
		Id:         0,
		Uuid:       uniqueid.NewUUID(),
		Content:    content,
		VisitorId:  0,
		KefuId:     0,
		MsgType:    publicType,
		Status:     0,
		IsRead:     0,
		IsSend:     0,
		SendTime:   0,
		ExpireTime: 0,
	}
}

func NewSystemMessage(content any) *types.Message {
	return &types.Message{
		Id:         0,
		Uuid:       uniqueid.NewUUID(),
		Content:    fmt.Sprintf("%s", content),
		VisitorId:  0,
		KefuId:     0,
		MsgType:    systemType,
		Status:     0,
		IsRead:     0,
		IsSend:     0,
		SendTime:   0,
		ExpireTime: 0,
	}
}
