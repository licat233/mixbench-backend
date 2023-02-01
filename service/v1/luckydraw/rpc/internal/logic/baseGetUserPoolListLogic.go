/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:37:47
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetUserPoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetUserPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetUserPoolListLogic {
	return &BaseGetUserPoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetUserPoolListLogic) BaseGetUserPoolList(in *luckydraw_pb.GetUserPoolListReq) (*luckydraw_pb.GetUserPoolListResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	list, total, err := l.svcCtx.UserPoolModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, tools.PbUserPoolToMdUserPool(in.UserPool))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetUserPoolListResp{
		UserPools: tools.MdUserPoolsToPbUserPools(list),
		Total:     total,
	}, nil
}
