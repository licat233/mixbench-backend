/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:39:15
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

type BasePutUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutUserLogic {
	return &BasePutUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BasePutUserLogic) BasePutUser(in *luckydraw_pb.PutUserReq) (*luckydraw_pb.PutUserResp, error) {
	data := &model.User{
		Id:       in.Id,
		Uuid:     in.Uuid,
		Passport: in.Passport,
		Email:    in.Email,
		Phone:    in.Phone,
		Name:     in.Name,
		Address:  in.Address,
		Score:    in.Score,
		Status:   in.Status,
	}

	if err := l.svcCtx.UserModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutUserResp{}, nil
}
