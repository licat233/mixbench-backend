package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/luckdraw/rpc/internal/svc"
	"mixbench/service/v1/luckdraw/rpc/luckdraw_pb"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckerListLogic {
	return &GetLuckerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询某用户在某活动里的中奖名单
func (l *GetLuckerListLogic) GetLuckerList(in *luckdraw_pb.GetLuckerListReq) (*luckdraw_pb.GetLuckerListResp, error) {
	//这里只能查询一个用户，所以user_id不能等于-1
	if in.UserId == -1 || in.UserId == 0 {
		return nil, errorx.ExtRpcErr(errors.New("userId invalid"))
	}
	activity, err := l.svcCtx.ActivityModel.FindOneByUuid(l.ctx, in.ActivityUuid)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("activity not found"))
	}

	options := model.NewLuckerDetailsOptions()
	options.Lucker.UserId = in.UserId
	options.ActivityId = activity.Id

	luckerDtailsList, _, err := l.svcCtx.CommonModel.FindsLuckerDetails(l.ctx, options)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckdraw_pb.GetLuckerListResp{
		LuckerDetails: convertLuckerDetailsList(luckerDtailsList),
	}, nil
}
