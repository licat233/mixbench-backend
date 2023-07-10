package dataconv

import (
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"
)

func MdChatRecordToApiChatRecord(md *model.ChatRecord) *types.ChatRecord {
	if md == nil {
		return nil
	}
	return &types.ChatRecord{
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
	}
}

func MdChatRecords2ApiChatRecords(mdList []*model.ChatRecord) []*types.ChatRecord {
	if mdList == nil {
		return nil
	}
	data := make([]*types.ChatRecord, 0)
	for _, v := range mdList {
		data = append(data, MdChatRecordToApiChatRecord(v))
	}
	return data
}

func MdIpblackToApiIpblack(md *model.Ipblack) *types.Ipblack {
	if md == nil {
		return nil
	}
	return &types.Ipblack{
		Id: md.Id,
		Ip: md.Ip,
	}
}

func MdIpblacks2ApiIpblacks(mdList []*model.Ipblack) []*types.Ipblack {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Ipblack, 0)
	for _, v := range mdList {
		data = append(data, MdIpblackToApiIpblack(v))
	}
	return data
}

func MdKefuToApiKefu(md *model.Kefu) *types.Kefu {
	if md == nil {
		return nil
	}
	return &types.Kefu{
		Id:        md.Id,
		Uuid:      md.Uuid,
		ProjectId: md.ProjectId,
		Passport:  md.Passport,
		Username:  md.Username,
		Password:  md.Password,
		Status:    md.Status,
	}
}

func MdKefus2ApiKefus(mdList []*model.Kefu) []*types.Kefu {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Kefu, 0)
	for _, v := range mdList {
		data = append(data, MdKefuToApiKefu(v))
	}
	return data
}

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

func MdProjectToApiProject(md *model.Project) *types.Project {
	if md == nil {
		return nil
	}
	return &types.Project{
		Id:     md.Id,
		Uuid:   md.Uuid,
		Name:   md.Name,
		Status: md.Status,
	}
}

func MdProjects2ApiProjects(mdList []*model.Project) []*types.Project {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Project, 0)
	for _, v := range mdList {
		data = append(data, MdProjectToApiProject(v))
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
