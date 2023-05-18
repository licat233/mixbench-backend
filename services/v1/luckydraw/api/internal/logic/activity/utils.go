package activity

import (
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"
	"time"
)

func TimeHandler(timeStr string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(timeLayout, timeStr, time.Local)
	if err != nil {
		return -1
	}
	return t.Local().Unix()
}

func PbActivityDetailToApiActivityDetail(data *luckydraw_pb.ActivityDetail) *types.ActivityDetail {
	if data == nil {
		return nil
	}
	res := &types.ActivityDetail{
		Activity:       *dataconv.PbActivityToApiActivity(data.Activity),
		ActivityPrizes: []types.ActivityPrize{},
	}

	for _, v := range data.ActivityPrizes {
		item := types.ActivityPrize{
			Prize:      *dataconv.PbPrizeToApiPrize(v.Prize),
			PrizePool:  *dataconv.PbPrizePoolToApiPrizePool(v.PrizePool),
			PrizeGrade: *dataconv.PbPrizeGradeToApiPrizeGrade(v.PrizeGrade),
		}
		res.ActivityPrizes = append(res.ActivityPrizes, item)
	}
	return res
}
