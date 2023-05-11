/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:25:38
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type SelectPrizesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPrizesLogic {
	return &SelectPrizesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectPrizesLogic) SelectPrizes(req *types.SelectPrizesReq) (resp any, err error) {
	if !govalidator.IsUUID(req.ActivityUUID) {
		return nil, errorx.ErrInvalidUuid
	}
	userId := ctxdata.GetUidFromCtx(l.ctx)
	in := &luckydraw_pb.SelectPrizesReq{
		UserId:       userId,
		ActivityId:   req.ActivityId,
		ActivityUuid: req.ActivityUUID,
		LuckyUuids:   req.LuckyUuidList,
	}
	res, err := l.svcCtx.LuckydrawRpc.SelectPrizes(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.Prizes, err)
}
