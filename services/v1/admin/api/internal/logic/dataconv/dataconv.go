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
