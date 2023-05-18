package project

import (
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"
)

func PbProjectDetailToApiProjectDetail(in *luckydraw_pb.ProjectDetail) *types.ProjectDtail {
	data := &types.ProjectDtail{
		Project:    *dataconv.PbProjectToApiProject(in.Project),
		Activities: []types.Activity{},
	}
	for _, v := range in.Activities {
		data.Activities = append(data.Activities, *dataconv.PbActivityToApiActivity(v))
	}
	return data
}
