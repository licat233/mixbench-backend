/*
 * @Author: licat
 * @Date: 2023-02-01 15:34:22
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:34:57
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/common/utils"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type BaseAddLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseAddLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseAddLuckyLogic {
	return &BaseAddLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseAddLuckyLogic) BaseAddLucky(in *luckydraw_pb.AddLuckyReq) (*luckydraw_pb.AddLuckyResp, error) {
	if in.Ip != "" && !govalidator.IsIP(in.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}

	lukeDate := utils.TimestampHandle(in.LuckDate)
	data := &model.Lucky{
		Uuid:        uniqueid.NewUUID(),
		UserId:      in.UserId,
		PrizePoolId: in.PrizePoolId,
		Status:      in.Status,
		LuckDate:    utils.Int64ToTime(lukeDate),
		Ip:          in.Ip,
	}
	result, err := l.svcCtx.LuckyModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	lucky, err := l.svcCtx.LuckyModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddLuckyResp{
		Lucky: tools.MdLuckyToPbLucky(lucky),
	}, nil
}
