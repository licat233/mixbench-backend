package dataconv

import (
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
)

func PbEnumToApiEnum(in *authentication_pb.Enum) *types.Enum {
	if in == nil {
		return nil
	}
	return &types.Enum{
		Label: in.Label,
		Value: in.Value,
	}
}

func PbEnumsToApiEnums(list []*authentication_pb.Enum) []*types.Enum {
	res := []*types.Enum{}
	for _, v := range list {
		res = append(res, PbEnumToApiEnum(v))
	}
	return res
}

func PbOptionToApiOption(in *authentication_pb.Option) *types.Option {
	if in == nil {
		return nil
	}
	return &types.Option{
		Title: in.Title,
		Value: in.Value,
	}
}

func PbOptionsToApiOptions(list []*authentication_pb.Option) []*types.Option {
	res := []*types.Option{}
	for _, v := range list {
		res = append(res, PbOptionToApiOption(v))
	}
	return res
}

func PbTreeOptionToApiTreeOption(in *authentication_pb.TreeOption) *types.TreeOption {
	if in == nil {
		return nil
	}
	return &types.TreeOption{
		Title:    in.Title,
		Value:    in.Value,
		Children: PbTreeOptionsToApiTreeOptions(in.Children),
	}
}

func PbTreeOptionsToApiTreeOptions(list []*authentication_pb.TreeOption) []types.TreeOption {
	res := []types.TreeOption{}
	for _, v := range list {
		res = append(res, *PbTreeOptionToApiTreeOption(v))
	}
	return res
}

func MdIpBlacklistToApiIpBlacklist(md *model.IpBlacklist) *types.IpBlacklist {
	if md == nil {
		return nil
	}
	return &types.IpBlacklist{
		Id: md.Id,
		Ip: md.Ip,
	}
}

func MdIpBlacklists2ApiIpBlacklists(mdList []*model.IpBlacklist) []*types.IpBlacklist {
	if mdList == nil {
		return nil
	}
	data := make([]*types.IpBlacklist, 0)
	for _, v := range mdList {
		data = append(data, MdIpBlacklistToApiIpBlacklist(v))
	}
	return data
}

func PbIpBlacklistToApiIpBlacklist(pb *authentication_pb.IpBlacklist) *types.IpBlacklist {
	if pb == nil {
		return nil
	}
	return &types.IpBlacklist{
		Id: pb.Id,
		Ip: pb.Ip,
	}
}

func PbIpBlacklists2ApiIpBlacklists(pbList []*authentication_pb.IpBlacklist) []*types.IpBlacklist {
	if pbList == nil {
		return nil
	}
	data := make([]*types.IpBlacklist, 0)
	for _, v := range pbList {
		data = append(data, PbIpBlacklistToApiIpBlacklist(v))
	}
	return data
}

func MdTokenBlacklistToApiTokenBlacklist(md *model.TokenBlacklist) *types.TokenBlacklist {
	if md == nil {
		return nil
	}
	return &types.TokenBlacklist{
		Id:       md.Id,
		Token:    md.Token,
		Platform: md.Platform,
		Ip:       md.Ip,
	}
}

func MdTokenBlacklists2ApiTokenBlacklists(mdList []*model.TokenBlacklist) []*types.TokenBlacklist {
	if mdList == nil {
		return nil
	}
	data := make([]*types.TokenBlacklist, 0)
	for _, v := range mdList {
		data = append(data, MdTokenBlacklistToApiTokenBlacklist(v))
	}
	return data
}

func PbTokenBlacklistToApiTokenBlacklist(pb *authentication_pb.TokenBlacklist) *types.TokenBlacklist {
	if pb == nil {
		return nil
	}
	return &types.TokenBlacklist{
		Id:       pb.Id,
		Token:    pb.Token,
		Platform: pb.Platform,
		Ip:       pb.Ip,
	}
}

func PbTokenBlacklists2ApiTokenBlacklists(pbList []*authentication_pb.TokenBlacklist) []*types.TokenBlacklist {
	if pbList == nil {
		return nil
	}
	data := make([]*types.TokenBlacklist, 0)
	for _, v := range pbList {
		data = append(data, PbTokenBlacklistToApiTokenBlacklist(v))
	}
	return data
}
