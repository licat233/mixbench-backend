package dataconv

import (
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
)

func PbEnumToApiEnum(in *admin_pb.Enum) *types.Enum {
	if in == nil {
		return nil
	}
	return &types.Enum{
		Label: in.Label,
		Value: in.Value,
	}
}

func PbEnumsToApiEnums(list []*admin_pb.Enum) []*types.Enum {
	res := []*types.Enum{}
	for _, v := range list {
		res = append(res, PbEnumToApiEnum(v))
	}
	return res
}

func PbOptionToApiOption(in *admin_pb.Option) *types.Option {
	if in == nil {
		return nil
	}
	return &types.Option{
		Title: in.Title,
		Value: in.Value,
	}
}

func PbOptionsToApiOptions(list []*admin_pb.Option) []*types.Option {
	res := []*types.Option{}
	for _, v := range list {
		res = append(res, PbOptionToApiOption(v))
	}
	return res
}

func PbTreeOptionToApiTreeOption(in *admin_pb.TreeOption) *types.TreeOption {
	if in == nil {
		return nil
	}
	return &types.TreeOption{
		Title:    in.Title,
		Value:    in.Value,
		Children: PbTreeOptionsToApiTreeOptions(in.Children),
	}
}

func PbTreeOptionsToApiTreeOptions(list []*admin_pb.TreeOption) []types.TreeOption {
	res := []types.TreeOption{}
	for _, v := range list {
		res = append(res, *PbTreeOptionToApiTreeOption(v))
	}
	return res
}

func MdAdminerToApiAdminer(md *model.Adminer) *types.Adminer {
	if md == nil {
		return nil
	}
	return &types.Adminer{
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

func MdAdminers2ApiAdminers(mdList []*model.Adminer) []*types.Adminer {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Adminer, 0)
	for _, v := range mdList {
		data = append(data, MdAdminerToApiAdminer(v))
	}
	return data
}

func PbAdminerToApiAdminer(pb *admin_pb.Adminer) *types.Adminer {
	if pb == nil {
		return nil
	}
	return &types.Adminer{
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
		LastLogin:    pb.LastLogin,
	}
}

func PbAdminers2ApiAdminers(pbList []*admin_pb.Adminer) []*types.Adminer {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Adminer, 0)
	for _, v := range pbList {
		data = append(data, PbAdminerToApiAdminer(v))
	}
	return data
}

func MdCampaignToApiCampaign(md *model.Campaign) *types.Campaign {
	if md == nil {
		return nil
	}
	return &types.Campaign{
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

func MdCampaigns2ApiCampaigns(mdList []*model.Campaign) []*types.Campaign {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Campaign, 0)
	for _, v := range mdList {
		data = append(data, MdCampaignToApiCampaign(v))
	}
	return data
}

func PbCampaignToApiCampaign(pb *admin_pb.Campaign) *types.Campaign {
	if pb == nil {
		return nil
	}
	return &types.Campaign{
		Id:        pb.Id,
		Uuid:      pb.Uuid,
		Name:      pb.Name,
		ProjectId: pb.ProjectId,
		Status:    pb.Status,
		Pattern:   pb.Pattern,
		MaxCount:  pb.MaxCount,
		MaxPrizes: pb.MaxPrizes,
		StartTime: pb.StartTime,
		EndTime:   pb.EndTime,
	}
}

func PbCampaigns2ApiCampaigns(pbList []*admin_pb.Campaign) []*types.Campaign {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Campaign, 0)
	for _, v := range pbList {
		data = append(data, PbCampaignToApiCampaign(v))
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

func PbProjectToApiProject(pb *admin_pb.Project) *types.Project {
	if pb == nil {
		return nil
	}
	return &types.Project{
		Id:     pb.Id,
		Uuid:   pb.Uuid,
		Name:   pb.Name,
		Status: pb.Status,
	}
}

func PbProjects2ApiProjects(pbList []*admin_pb.Project) []*types.Project {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Project, 0)
	for _, v := range pbList {
		data = append(data, PbProjectToApiProject(v))
	}
	return data
}

func MdWebToApiWeb(md *model.Web) *types.Web {
	if md == nil {
		return nil
	}
	return &types.Web{
		Id:          md.Id,
		Uuid:        md.Uuid,
		ProjectId:   md.ProjectId,
		Url:         md.Url,
		Host:        md.Host,
		Description: md.Description,
		Code:        md.Code,
	}
}

func MdWebs2ApiWebs(mdList []*model.Web) []*types.Web {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Web, 0)
	for _, v := range mdList {
		data = append(data, MdWebToApiWeb(v))
	}
	return data
}

func PbWebToApiWeb(pb *admin_pb.Web) *types.Web {
	if pb == nil {
		return nil
	}
	return &types.Web{
		Id:          pb.Id,
		Uuid:        pb.Uuid,
		ProjectId:   pb.ProjectId,
		Url:         pb.Url,
		Host:        pb.Host,
		Description: pb.Description,
		Code:        pb.Code,
	}
}

func PbWebs2ApiWebs(pbList []*admin_pb.Web) []*types.Web {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Web, 0)
	for _, v := range pbList {
		data = append(data, PbWebToApiWeb(v))
	}
	return data
}
