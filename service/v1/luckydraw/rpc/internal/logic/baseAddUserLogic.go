/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:39:56
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseAddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseAddUserLogic {
	return &BaseAddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseAddUserLogic) BaseAddUser(in *luckydraw_pb.AddUserReq) (*luckydraw_pb.AddUserResp, error) {
	data := &model.User{
		Uuid:     uniqueid.NewUUID(),
		Passport: in.Passport,
		Email:    in.Email,
		Phone:    in.Phone,
		Name:     in.Name,
		Address:  in.Address,
		Score:    in.Score,
		Status:   in.Status,
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddUserResp{
		User: tools.MdUserToPbUser(user),
	}, nil
}
