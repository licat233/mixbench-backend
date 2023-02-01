/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:38:27
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type BasePutLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutLuckyLogic {
	return &BasePutLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BasePutLuckyLogic) BasePutLucky(in *luckydraw_pb.PutLuckyReq) (*luckydraw_pb.PutLuckyResp, error) {
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

	return &luckydraw_pb.PutLuckyResp{}, nil
}
