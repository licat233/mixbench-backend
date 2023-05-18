package luckydrawlogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckyDetailsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckyDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyDetailsListLogic {
	return &GetLuckyDetailsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取中奖详情名单
func (l *GetLuckyDetailsListLogic) GetLuckyDetailsList(in *luckydraw_pb.LuckyDetailsReq) (*luckydraw_pb.GetLuckyDetailsResp, error) {
	if in.Lucky != nil && in.Lucky.Ip != "" && !govalidator.IsIP(in.Lucky.Ip) {
		return &luckydraw_pb.GetLuckyDetailsResp{
			Data:  []*luckydraw_pb.LuckyDetail{},
			Total: 0,
		}, nil
	}
	list, tatol, err := l.svcCtx.CommonModel.FindsLuckyDetails(l.ctx, &model.LuckyDetailOptions{
		PageSize:     in.ListReq.PageSize,
		Page:         in.ListReq.Page,
		Keyword:      in.ListReq.Keyword,
		Lucky:        &model.Lucky{UserId: in.Lucky.UserId, PrizePoolId: in.Lucky.PrizePoolId, Status: in.Lucky.Status, LuckDate: utils.Int64ToTime(in.Lucky.LuckDate), Ip: in.Lucky.Ip},
		UserName:     in.UserName,
		PrizeName:    in.PrizeName,
		ProjectId:    in.ProjectId,
		ActivityId:   in.ActivityId,
		ActivityName: in.ActivityName,
		PrizeId:      in.PrizeId,
		PrizeGradeId: in.PrizeGradeId,
	})

	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.GetLuckyDetailsResp{
		Data:  convertLuckyDetailsList(list),
		Total: tatol,
	}, nil
}

func convertLuckyDetailsList(in []*model.LuckyDetail) []*luckydraw_pb.LuckyDetail {
	var res []*luckydraw_pb.LuckyDetail
	for _, item := range in {
		res = append(res, &luckydraw_pb.LuckyDetail{
			Lucky:      dataconv.MdLuckyToPbLucky(item.Lucky),
			User:       dataconv.MdUserToPbUser(item.User),
			PrizePool:  dataconv.MdPrizePoolToPbPrizePool(item.PrizePool),
			Activity:   dataconv.MdActivityToPbActivity(item.Activity),
			Prize:      dataconv.MdPrizeToPbPrize(item.Prize),
			PrizeGrade: dataconv.MdPrizeGradeToPbPrizeGrade(item.PrizeGrade),
			Project:    dataconv.MdProjectToPbProject(item.Project),
		})
	}
	return res
}
