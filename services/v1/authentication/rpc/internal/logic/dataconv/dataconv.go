package dataconv

import (
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
)

func PbIpblackToMdIpblack(pb *authentication_pb.Ipblack) *model.Ipblack {
	if pb == nil {
		return nil
	}
	return &model.Ipblack{
		Id: pb.Id,
		Ip: pb.Ip,
	}
}

func MdIpblackToPbIpblack(md *model.Ipblack) *authentication_pb.Ipblack {
	if md == nil {
		return nil
	}
	return &authentication_pb.Ipblack{
		Id: md.Id,
		Ip: md.Ip,
	}
}

func PbIpblacks2MdIpblacks(pbList []*authentication_pb.Ipblack) []*model.Ipblack {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Ipblack, 0)
	for _, v := range pbList {
		data = append(data, PbIpblackToMdIpblack(v))
	}
	return data
}

func MdIpblacks2PbIpblacks(mdList []*model.Ipblack) []*authentication_pb.Ipblack {
	if mdList == nil {
		return nil
	}
	data := make([]*authentication_pb.Ipblack, 0)
	for _, v := range mdList {
		data = append(data, MdIpblackToPbIpblack(v))
	}
	return data
}

func ListReqParams(req *authentication_pb.ListReq) (pageSize int64, page int64, keyword string) {
	if req != nil {
		pageSize = req.PageSize
		page = req.Page
		keyword = req.Keyword
	}
	return
}
