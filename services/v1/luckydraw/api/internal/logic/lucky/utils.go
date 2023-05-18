package lucky

import (
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"
)

func PbLuckyDetailToApiLuckyDetail(data *luckydraw_pb.LuckyDetail) *types.LuckyDetail {
	if data == nil {
		return nil
	}
	res := &types.LuckyDetail{
		Lucky:      *dataconv.PbLuckyToApiLucky(data.Lucky),
		User:       *dataconv.PbUserToApiUser(data.User),
		PrizePool:  *dataconv.PbPrizePoolToApiPrizePool(data.PrizePool),
		Prize:      *dataconv.PbPrizeToApiPrize(data.Prize),
		PrizeGrade: *dataconv.PbPrizeGradeToApiPrizeGrade(data.PrizeGrade),
		Activity:   *dataconv.PbActivityToApiActivity(data.Activity),
		Project:    *dataconv.PbProjectToApiProject(data.Project),
	}
	return res
}
