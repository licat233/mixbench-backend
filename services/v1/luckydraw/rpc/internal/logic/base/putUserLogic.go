package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUserLogic {
	return &PutUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户名单
func (l *PutUserLogic) PutUser(in *luckydraw_pb.PutUserReq) (*luckydraw_pb.PutUserResp, error) {
	data := &model.User{
		Id:       in.Id,
		Uuid:     in.Uuid,
		Passport: in.Passport,
		Email:    in.Email,
		Phone:    in.Phone,
		Name:     in.Name,
		Address:  in.Address,
		Score:    in.Score,
		Status:   in.Status,
	}

	if err := l.svcCtx.UserModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutUserResp{}, nil
}
