package types

import "mixbench/services/v1/chat/model"

var (
	TaskTypeSendMessage = "handle_send_message"
)

type SendMessageReq struct {
	*model.Message
}
