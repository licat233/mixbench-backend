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

func ListReqParams(req *admin_pb.ListReq) (pageSize int64, page int64, keyword string) {
	if req != nil {
		pageSize = req.PageSize
		page = req.Page
		keyword = req.Keyword
	}
	return
}
