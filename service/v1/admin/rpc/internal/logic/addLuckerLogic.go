/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 11:57:02
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/common/utils"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLuckyLogic {
	return &AddLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLuckyLogic) AddLucky(in *admin_pb.Lucky) (*admin_pb.AddLuckyResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}
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

	return &admin_pb.AddLuckyResp{
		Lucky: mdLuckyToPbLucky(lucky),
	}, nil
}
