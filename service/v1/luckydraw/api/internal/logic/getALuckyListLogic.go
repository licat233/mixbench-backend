/*
 * @Author: licat
 * @Date: 2023-02-20 18:10:59
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:26:00
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetALuckyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetALuckyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetALuckyListLogic {
	return &GetALuckyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetALuckyListLogic) GetALuckyList(req *types.GetALuckyListReq) (any, error) {
	//获取用户id
	userId := ctxdata.GetUidFromCtx(l.ctx)
	//活动uuid
	if !govalidator.IsUUID(req.ActivityUUID) {
		return nil, errorx.ErrInvalidUuid
	}
	in := &luckydraw_pb.GetUALuckyDetailsReq{
		UserId:       userId,
		ActivityUuid: req.ActivityUUID,
		ActivityId:   0,
	}
	res, err := l.svcCtx.LuckydrawRpc.LuckydrawClient.GetUALuckyDetails(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.LuckyDetails, err)
}
