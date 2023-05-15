/*
 * @Author: licat
 * @Date: 2023-02-20 18:10:59
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:11:20
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIdentifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIdentifyLogic {
	return &UserIdentifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserIdentifyLogic) UserIdentify(req *types.UserIdentifyReq) (any, error) {
	in := &luckydraw_pb.UserIdentifyReq{
		Passport: req.Passport,
	}
	resp, err := l.svcCtx.LuckydrawRpc.UserIdentify(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.JwtToken, err)
}
