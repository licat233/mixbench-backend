/*
 * @Author: licat
 * @Date: 2023-01-11 15:03:41
 * @LastEditors: licat
 * @LastEditTime: 2023-01-12 10:49:58
 * @Description: licat233@gmail.com
 */
package goods

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsAddLogic {
	return &GoodsAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsAddLogic) GoodsAdd(req *types.GoodsAddReq) (resp *types.BaseResp, err error) {

	return
}
