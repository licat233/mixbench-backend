package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityLogic {
	return &AddActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityLogic) AddActivity(req *types.AddActivityReq) (any, error) {
	in := &luckydraw_pb.AddActivityReq{
		Name:      req.Name,
		ProjectId: req.ProjectId,
		Status:    req.Status,
		Pattern:   req.Pattern,
		MaxCount:  req.MaxCount,
		MaxPrizes: req.MaxPrizes,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	_, err := l.svcCtx.LuckydrawRpc.AddActivity(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
