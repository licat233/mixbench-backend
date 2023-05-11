package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutActivityLogic {
	return &PutActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutActivityLogic) PutActivity(req *types.PutActivityReq) (any, error) {
	in := &luckydraw_pb.PutActivityReq{
		Id:        req.Id,
		Uuid:      req.Uuid,
		Name:      req.Name,
		ProjectId: req.ProjectId,
		Status:    req.Status,
		Pattern:   req.Pattern,
		MaxCount:  req.MaxCount,
		MaxPrizes: req.MaxPrizes,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	_, err := l.svcCtx.LuckydrawRpc.PutActivity(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
