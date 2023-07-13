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

func PbIpblackToApiIpblack(pb *authentication_pb.Ipblack) *types.Ipblack {
	if pb == nil {
		return nil
	}
	return &types.Ipblack{
		Id: pb.Id,
		Ip: pb.Ip,
	}
}

func PbIpblacks2ApiIpblacks(pbList []*authentication_pb.Ipblack) []*types.Ipblack {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Ipblack, 0)
	for _, v := range pbList {
		data = append(data, PbIpblackToApiIpblack(v))
	}
	return data
}
