package store

import (
	"mixbench/services/v1/chat/model"
	"time"
)

type RecentMessage struct {
	messages []*model.Message
	lastTime time.Time
}

func NewRecentMessage() *RecentMessage {
	return &RecentMessage{
		messages: make([]*model.Message, 0, 5),
		lastTime: time.Now(),
	}
}

func (m *RecentMessage) push(message *model.Message) {
	// 向切片中添加值
	// 如果切片长度已达到容量上限，则删除第一个值
	if len(m.messages) >= 5 {
		m.messages = m.messages[1:]
	}
	m.messages = append(m.messages, message)
}

func (m *RecentMessage) Validator(message *model.Message) bool {
	if message.SendTime.Sub(m.lastTime) < time.Millisecond*500 {
		//发送频率太快，小于500毫秒
		return false
	}
	m.lastTime = time.Now()
	m.push(message)

	//检查连续5条的信息内容是否都相同
	//检查连续5条的信息时间间隔是否都小于1s
	if len(m.messages) == 5 {
		contentStatus := false
		timeStatus := false
		for index, msg := range m.messages {
			if msg.Content != message.Content && !contentStatus {
				contentStatus = true
			}
			if index > 0 && !timeStatus {
				lasgMsg := m.messages[index-1]
				if msg.SendTime.Sub(lasgMsg.SendTime) >= time.Second {
					//如果存在时间间隔大于1s的，表示通过验证
					timeStatus = true
				}
			}
		}
		return contentStatus && timeStatus
	}

	return true
}
