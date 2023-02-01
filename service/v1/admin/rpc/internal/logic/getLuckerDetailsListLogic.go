/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 14:00:57
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

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

func (l *GetLuckyDetailsListLogic) GetLuckyDetailsList(in *admin_pb.LuckyDetailsReq) (*admin_pb.GetLuckyDetailsResp, error) {
	if in.Lucky != nil && in.Lucky.Ip != "" && !govalidator.IsIP(in.Lucky.Ip) {
		return &admin_pb.GetLuckyDetailsResp{
			Data:  []*admin_pb.LuckyDetail{},
			Total: 0,
		}, nil
	}
	list, tatol, err := l.svcCtx.CommonModel.FindsLuckyDetails(l.ctx, &model.LuckyDetailsOptions{
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
	return &admin_pb.GetLuckyDetailsResp{
		Data:  convertLuckyDetailsList(list),
		Total: tatol,
	}, nil
}
