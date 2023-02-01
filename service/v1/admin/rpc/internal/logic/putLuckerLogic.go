package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"
	"mixbench/common/utils"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type PutLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLuckyLogic {
	return &PutLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutLuckyLogic) PutLucky(in *admin_pb.LuckyPutReq) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}
	if in.Ip != "" && !govalidator.IsIP(in.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}
	lukeDate := utils.TimestampHandle(in.LuckDate)
	data := &model.Lucky{
		Id:          in.Id,
		Uuid:        in.Uuid,
		UserId:      in.UserId,
		PrizePoolId: in.PrizePoolId,
		Status:      in.Status,
		LuckDate:    utils.Int64ToTime(lukeDate),
		Ip:          in.Ip,
	}

	if err := l.svcCtx.LuckyModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
