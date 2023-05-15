package dataconv

import (
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
)

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
