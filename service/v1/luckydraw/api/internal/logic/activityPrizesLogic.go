/*
 * @Author: licat
 * @Date: 2023-02-20 18:10:59
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:26:07
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPrizesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPrizesLogic {
	return &ActivityPrizesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityPrizesLogic) ActivityPrizes(req *types.ActivityPrizesReq) (any, error) {
	if !govalidator.IsUUID(req.ActivityUUID) {
		return nil, errorx.ErrInvalidUuid
	}
	in := &luckydraw_pb.GetActivityPrizesReq{
		ActivityUuid: req.ActivityUUID,
		ActivityId:   0,
	}
	res, err := l.svcCtx.LuckydrawRpc.GetActivityPrizes(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.ActivityPrizes, err)
}
