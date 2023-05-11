/*
 * @Author: licat
 * @Date: 2023-02-20 18:10:59
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:24:06
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLuckyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLuckyLogic {
	return &SearchLuckyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 根据passport查询用户的中奖信息
func (l *SearchLuckyLogic) SearchLucky(req *types.SearchLuckyReq) (any, error) {
	in := &luckydraw_pb.GetUserLuckyDetailReq{
		Passport: req.Keyword,
	}
	res, err := l.svcCtx.LuckydrawRpc.GetUserLuckyDetail(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res, err)
}
