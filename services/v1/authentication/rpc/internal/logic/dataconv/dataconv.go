package dataconv

import (
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
)

func PbIpBlacklistToMdIpBlacklist(pb *authentication_pb.IpBlacklist) *model.IpBlacklist {
	if pb == nil {
		return nil
	}
	return &model.IpBlacklist{
		Id: pb.Id,
		Ip: pb.Ip,
	}
}

func MdIpBlacklistToPbIpBlacklist(md *model.IpBlacklist) *authentication_pb.IpBlacklist {
	if md == nil {
		return nil
	}
	return &authentication_pb.IpBlacklist{
		Id: md.Id,
		Ip: md.Ip,
	}
}

func PbIpBlacklists2MdIpBlacklists(pbList []*authentication_pb.IpBlacklist) []*model.IpBlacklist {
	if pbList == nil {
		return nil
	}
	data := make([]*model.IpBlacklist, 0)
	for _, v := range pbList {
		data = append(data, PbIpBlacklistToMdIpBlacklist(v))
	}
	return data
}

func MdIpBlacklists2PbIpBlacklists(mdList []*model.IpBlacklist) []*authentication_pb.IpBlacklist {
	if mdList == nil {
		return nil
	}
	data := make([]*authentication_pb.IpBlacklist, 0)
	for _, v := range mdList {
		data = append(data, MdIpBlacklistToPbIpBlacklist(v))
	}
	return data
}

func PbTokenBlacklistToMdTokenBlacklist(pb *authentication_pb.TokenBlacklist) *model.TokenBlacklist {
	if pb == nil {
		return nil
	}
	return &model.TokenBlacklist{
		Id:       pb.Id,
		Token:    pb.Token,
		Platform: pb.Platform,
		Ip:       pb.Ip,
	}
}

func MdTokenBlacklistToPbTokenBlacklist(md *model.TokenBlacklist) *authentication_pb.TokenBlacklist {
	if md == nil {
		return nil
	}
	return &authentication_pb.TokenBlacklist{
		Id:       md.Id,
		Token:    md.Token,
		Platform: md.Platform,
		Ip:       md.Ip,
	}
}

func PbTokenBlacklists2MdTokenBlacklists(pbList []*authentication_pb.TokenBlacklist) []*model.TokenBlacklist {
	if pbList == nil {
		return nil
	}
	data := make([]*model.TokenBlacklist, 0)
	for _, v := range pbList {
		data = append(data, PbTokenBlacklistToMdTokenBlacklist(v))
	}
	return data
}

func MdTokenBlacklists2PbTokenBlacklists(mdList []*model.TokenBlacklist) []*authentication_pb.TokenBlacklist {
	if mdList == nil {
		return nil
	}
	data := make([]*authentication_pb.TokenBlacklist, 0)
	for _, v := range mdList {
		data = append(data, MdTokenBlacklistToPbTokenBlacklist(v))
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
