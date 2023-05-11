package user

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUserLogic {
	return &PutUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutUserLogic) PutUser(req *types.PutUserReq) (any, error) {
	in := &luckydraw_pb.PutUserReq{
		Id:       req.Id,
		Uuid:     req.Uuid,
		Passport: req.Passport,
		Email:    req.Email,
		Phone:    req.Phone,
		Name:     req.Name,
		Address:  req.Address,
		Score:    req.Score,
		Status:   req.Status,
	}
	_, err := l.svcCtx.LuckydrawRpc.PutUser(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
