package prize

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizePutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizePutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizePutLogic {
	return &PrizePutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizePutLogic) PrizePut(req *types.PrizePutReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutPrize(l.ctx, &admin_pb.Prize{
		Id:     req.Id,
		Uuid:   req.Uuid,
		Name:   req.Name,
		Image:  req.Image,
		Price:  req.Price,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
