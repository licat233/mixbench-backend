package dataconv

import (
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"time"
)

func PbAdminerToMdAdminer(pb *admin_pb.Adminer) *model.Adminer {
	if pb == nil {
		return nil
	}
	return &model.Adminer{
		Id:           pb.Id,
		Uuid:         pb.Uuid,
		Name:         pb.Name,
		Avatar:       pb.Avatar,
		Access:       pb.Access,
		Passport:     pb.Passport,
		Password:     pb.Password,
		Email:        pb.Email,
		Status:       pb.Status,
		IsSuperAdmin: pb.IsSuperAdmin,
		ProjectId:    pb.ProjectId,
		LoginCount:   pb.LoginCount,
		LastLogin:    time.Unix(pb.LastLogin, 0).Local(),
	}
}

func MdAdminerToPbAdminer(md *model.Adminer) *admin_pb.Adminer {
	if md == nil {
		return nil
	}
	return &admin_pb.Adminer{
		Id:           md.Id,
		Uuid:         md.Uuid,
		Name:         md.Name,
		Avatar:       md.Avatar,
		Access:       md.Access,
		Passport:     md.Passport,
		Password:     md.Password,
		Email:        md.Email,
		Status:       md.Status,
		IsSuperAdmin: md.IsSuperAdmin,
		ProjectId:    md.ProjectId,
		LoginCount:   md.LoginCount,
		LastLogin:    md.LastLogin.Unix(),
	}
}

func PbAdminers2MdAdminers(pbList []*admin_pb.Adminer) []*model.Adminer {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Adminer, 0)
	for _, v := range pbList {
		data = append(data, PbAdminerToMdAdminer(v))
	}
	return data
}

func MdAdminers2PbAdminers(mdList []*model.Adminer) []*admin_pb.Adminer {
	if mdList == nil {
		return nil
	}
	data := make([]*admin_pb.Adminer, 0)
	for _, v := range mdList {
		data = append(data, MdAdminerToPbAdminer(v))
	}
	return data
}

func PbCampaignToMdCampaign(pb *admin_pb.Campaign) *model.Campaign {
	if pb == nil {
		return nil
	}
	return &model.Campaign{
		Id:        pb.Id,
		Uuid:      pb.Uuid,
		Name:      pb.Name,
		ProjectId: pb.ProjectId,
		Status:    pb.Status,
		Pattern:   pb.Pattern,
		MaxCount:  pb.MaxCount,
		MaxPrizes: pb.MaxPrizes,
		StartTime: time.Unix(pb.StartTime, 0).Local(),
		EndTime:   time.Unix(pb.EndTime, 0).Local(),
	}
}

func MdCampaignToPbCampaign(md *model.Campaign) *admin_pb.Campaign {
	if md == nil {
		return nil
	}
	return &admin_pb.Campaign{
		Id:        md.Id,
		Uuid:      md.Uuid,
		Name:      md.Name,
		ProjectId: md.ProjectId,
		Status:    md.Status,
		Pattern:   md.Pattern,
		MaxCount:  md.MaxCount,
		MaxPrizes: md.MaxPrizes,
		StartTime: md.StartTime.Unix(),
		EndTime:   md.EndTime.Unix(),
	}
}

func PbCampaigns2MdCampaigns(pbList []*admin_pb.Campaign) []*model.Campaign {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Campaign, 0)
	for _, v := range pbList {
		data = append(data, PbCampaignToMdCampaign(v))
	}
	return data
}

func MdCampaigns2PbCampaigns(mdList []*model.Campaign) []*admin_pb.Campaign {
	if mdList == nil {
		return nil
	}
	data := make([]*admin_pb.Campaign, 0)
	for _, v := range mdList {
		data = append(data, MdCampaignToPbCampaign(v))
	}
	return data
}

func PbProjectToMdProject(pb *admin_pb.Project) *model.Project {
	if pb == nil {
		return nil
	}
	return &model.Project{
		Id:     pb.Id,
		Uuid:   pb.Uuid,
		Name:   pb.Name,
		Status: pb.Status,
	}
}

func MdProjectToPbProject(md *model.Project) *admin_pb.Project {
	if md == nil {
		return nil
	}
	return &admin_pb.Project{
		Id:     md.Id,
		Uuid:   md.Uuid,
		Name:   md.Name,
		Status: md.Status,
	}
}

func PbProjects2MdProjects(pbList []*admin_pb.Project) []*model.Project {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Project, 0)
	for _, v := range pbList {
		data = append(data, PbProjectToMdProject(v))
	}
	return data
}

func MdProjects2PbProjects(mdList []*model.Project) []*admin_pb.Project {
	if mdList == nil {
		return nil
	}
	data := make([]*admin_pb.Project, 0)
	for _, v := range mdList {
		data = append(data, MdProjectToPbProject(v))
	}
	return data
}

func PbWebToMdWeb(pb *admin_pb.Web) *model.Web {
	if pb == nil {
		return nil
	}
	return &model.Web{
		Id:          pb.Id,
		Uuid:        pb.Uuid,
		ProjectId:   pb.ProjectId,
		Url:         pb.Url,
		Host:        pb.Host,
		Description: pb.Description,
		Code:        pb.Code,
	}
}

func MdWebToPbWeb(md *model.Web) *admin_pb.Web {
	if md == nil {
		return nil
	}
	return &admin_pb.Web{
		Id:          md.Id,
		Uuid:        md.Uuid,
		ProjectId:   md.ProjectId,
		Url:         md.Url,
		Host:        md.Host,
		Description: md.Description,
		Code:        md.Code,
	}
}

func PbWebs2MdWebs(pbList []*admin_pb.Web) []*model.Web {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Web, 0)
	for _, v := range pbList {
		data = append(data, PbWebToMdWeb(v))
	}
	return data
}

func MdWebs2PbWebs(mdList []*model.Web) []*admin_pb.Web {
	if mdList == nil {
		return nil
	}
	data := make([]*admin_pb.Web, 0)
	for _, v := range mdList {
		data = append(data, MdWebToPbWeb(v))
	}
	return data
}

func ListReqParams(req *admin_pb.ListReq) (pageSize int64, page int64, keyword string) {
	if req != nil {
		pageSize = req.PageSize
		page = req.Page
		keyword = req.Keyword
	}
	return
}
