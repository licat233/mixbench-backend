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

type AddLuckerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLuckerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLuckerLogic {
	return &AddLuckerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLuckerLogic) AddLucker(in *admin_pb.Lucker) (*admin_pb.AddLuckerResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}
	if in.Ip != "" && !govalidator.IsIP(in.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}

	lukeDate := utils.TimestampHandle(in.LuckDate)
	data := &model.Lucker{
		Uuid:         uniqueid.NewUUID(),
		UserId:       in.UserId,
		PrizesPoolId: in.PrizesPoolId,
		Status:       in.Status,
		LuckDate:     utils.Int64ToTime(lukeDate),
		Ip:           in.Ip,
	}
	result, err := l.svcCtx.LuckerModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	lucker, err := l.svcCtx.LuckerModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddLuckerResp{
		Lucker: mdLuckerToPbLucker(lucker),
	}, nil
}
