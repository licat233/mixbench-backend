package user

import (
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"
)

func PbUserDetailToApiUserDetail(in *luckydraw_pb.UserDetail) *types.UserDetail {
	data := &types.UserDetail{
		User:           *dataconv.PbUserToApiUser(in.User),
		UserActivities: []types.UserActivity{},
	}
	for _, item := range in.UserActivities {
		ua := types.UserActivity{
			Activity: *dataconv.PbActivityToApiActivity(item.Activity),
			UserPool: *dataconv.PbUserPoolToApiUserPool(item.UserPool),
		}
		data.UserActivities = append(data.UserActivities, ua)
	}
	return data
}
