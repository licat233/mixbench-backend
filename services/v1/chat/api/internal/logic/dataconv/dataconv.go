package dataconv

import (
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"
)

func MdMessageToApiMessage(md *model.Message) *types.Message {
	if md == nil {
		return nil
	}
	return &types.Message{
		Id:          md.Id,
		Uuid:        md.Uuid,
		Content:     md.Content,
		ContentType: md.ContentType,
		VisitorId:   md.VisitorId,
		KefuId:      md.KefuId,
		ProjectId:   md.ProjectId,
		MsgType:     md.MsgType,
		Status:      md.Status,
		IsRead:      md.IsRead,
		IsSend:      md.IsSend,
		SendTime:    md.SendTime.Unix(),
		ExpireTime:  md.ExpireTime.Unix(),
	}
}

func MdMessages2ApiMessages(mdList []*model.Message) []*types.Message {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Message, 0)
	for _, v := range mdList {
		data = append(data, MdMessageToApiMessage(v))
	}
	return data
}

func MdMessageConfigToApiMessageConfig(md *model.MessageConfig) *types.MessageConfig {
	if md == nil {
		return nil
	}
	return &types.MessageConfig{
		Id:        md.Id,
		ProjectId: md.ProjectId,
		Name:      md.Name,
		Msgkey:    md.Msgkey,
		Content:   md.Content,
	}
}

func MdMessageConfigs2ApiMessageConfigs(mdList []*model.MessageConfig) []*types.MessageConfig {
	if mdList == nil {
		return nil
	}
	data := make([]*types.MessageConfig, 0)
	for _, v := range mdList {
		data = append(data, MdMessageConfigToApiMessageConfig(v))
	}
	return data
}

func MdVisitorToApiVisitor(md *model.Visitor) *types.Visitor {
	if md == nil {
		return nil
	}
	return &types.Visitor{
		Id:        md.Id,
		Uuid:      md.Uuid,
		Name:      md.Name,
		Ip:        md.Ip,
		KefuId:    md.KefuId,
		ProjectId: md.ProjectId,
		Status:    md.Status,
		Referer:   md.Referer,
		Region:    md.Region,
		City:      md.City,
		Platform:  md.Platform,
		Os:        md.Os,
		Browser:   md.Browser,
		Language:  md.Language,
	}
}

func MdVisitors2ApiVisitors(mdList []*model.Visitor) []*types.Visitor {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Visitor, 0)
	for _, v := range mdList {
		data = append(data, MdVisitorToApiVisitor(v))
	}
	return data
}
