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

type GetLuckerDetailsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckerDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckerDetailsListLogic {
	return &GetLuckerDetailsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckerDetailsListLogic) GetLuckerDetailsList(in *admin_pb.LuckerDetailsReq) (*admin_pb.GetLuckerDetailsResp, error) {
	if in.Lucker != nil && in.Lucker.Ip != "" && !govalidator.IsIP(in.Lucker.Ip) {
		return &admin_pb.GetLuckerDetailsResp{
			Data:  []*admin_pb.LuckerDetail{},
			Total: 0,
		}, nil
	}
	list, tatol, err := l.svcCtx.CommonModel.FindsLuckerDetails(l.ctx, &model.LuckerDetailsOptions{
		PageSize:     in.ListReq.PageSize,
		Page:         in.ListReq.Page,
		Keyword:      in.ListReq.Keyword,
		Lucker:       &model.Lucker{UserId: in.Lucker.UserId, PrizesPoolId: in.Lucker.PrizesPoolId, Status: in.Lucker.Status, LuckDate: utils.Int64ToTime(in.Lucker.LuckDate), Ip: in.Lucker.Ip},
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
	return &admin_pb.GetLuckerDetailsResp{
		Data:  convertLuckerDetailsList(list),
		Total: tatol,
	}, nil
}
