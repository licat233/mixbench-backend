package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPrizesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPrizesLogic {
	return &ActivityPrizesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityPrizesLogic) ActivityPrizes(in *admin_pb.ActivityPrizesReq) (*admin_pb.ActivityPrizesResp, error) {
	prizesPoolLogic := NewSearchPrizesPoolLogic(l.ctx, l.svcCtx)
	prizesPoolListResp, err := prizesPoolLogic.SearchPrizesPool(&admin_pb.SearchPrizesPoolReq{
		ActivityId: in.ActivityId,
	})
	if err != nil {
		return nil, err
	}
	var prizesIds []int64
	for _, prizesPool := range prizesPoolListResp.Data {
		prizesIds = append(prizesIds, prizesPool.PrizeId)
	}

	prizeList, err := l.svcCtx.PrizeModel.FindInIds(l.ctx, prizesIds)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}

	prizeGradeListLogic := NewGetPrizeGradeListLogic(l.ctx, l.svcCtx)
	prizeGradeListResp, err := prizeGradeListLogic.GetPrizeGradeList(&admin_pb.PrizeGradeListReq{
		ListReq:    nil,
		PrizeGrade: nil,
	})
	if err != nil {
		return nil, err
	}

	var activityPrizes []*admin_pb.ActivityPrize
	for _, prize := range prizeList {
		for _, prizesPool := range prizesPoolListResp.Data {
			if prize.Id == prizesPool.PrizeId {
				for _, prizeGrade := range prizeGradeListResp.Data {
					if prizesPool.GradeId == prizeGrade.Id {
						activityPrizes = append(activityPrizes, &admin_pb.ActivityPrize{
							Prize:      mdPrizeToPbPrize(prize),
							PrizesPool: prizesPool,
							PrizeGrade: prizeGrade,
						})
					}
				}
			}
		}
	}

	return &admin_pb.ActivityPrizesResp{
		Data: activityPrizes,
	}, nil
}
