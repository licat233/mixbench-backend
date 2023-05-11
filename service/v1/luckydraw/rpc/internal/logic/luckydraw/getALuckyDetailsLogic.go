/*
 * @Author: licat
 * @Date: 2023-02-01 15:34:22
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:41:58
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetALuckyDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetALuckyDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetALuckyDetailsLogic {
	return &GetALuckyDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取某活动的中奖详细列表
func (l *GetALuckyDetailsLogic) GetALuckyDetails(in *luckydraw_pb.GetALuckyDetailsReq) (*luckydraw_pb.GetALuckyDetailsResp, error) {
	if in.Lucky != nil && in.Lucky.Ip != "" && !govalidator.IsIP(in.Lucky.Ip) {
		return &luckydraw_pb.GetALuckyDetailsResp{
			LuckyDetails: []*luckydraw_pb.LuckyDetail{},
			Total:        0,
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
	return &luckydraw_pb.GetALuckyDetailsResp{
		LuckyDetails: tools.ConvertLuckyDetails(list),
		Total:        tatol,
	}, nil
}
