package wshub

import (
	"fmt"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/model"
)

func NewSystemMessage(content any, kefuId, visitorId int64) *model.Message {
	return &model.Message{
		Id:          0,
		Uuid:        uniqueid.NewUUID(),
		Content:     fmt.Sprintf("%s", content),
		ContentType: 0,
		VisitorId:   visitorId,
		KefuId:      kefuId,
		ProjectId:   0,
		MsgType:     SystemMsgType,
		Status:      0,
		IsRead:      0,
		IsSend:      0,
	}
}
