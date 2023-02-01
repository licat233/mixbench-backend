/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:38:44
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BasePutPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutPrizeLogic {
	return &BasePutPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BasePutPrizeLogic) BasePutPrize(in *luckydraw_pb.PutPrizeReq) (*luckydraw_pb.PutPrizeResp, error) {
	data := &model.Prize{
		Id:     in.Id,
		Uuid:   in.Uuid,
		Name:   in.Name,
		Image:  in.Image,
		Price:  in.Price,
		Status: in.Status,
	}

	if err := l.svcCtx.PrizeModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutPrizeResp{}, nil
}
