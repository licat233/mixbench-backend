/*
 * @Author: licat
 * @Date: 2023-02-01 16:56:01
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 17:31:46
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"
	"strings"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByKeywordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByKeywordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByKeywordLogic {
	return &GetUserByKeywordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过关键词查找用户【精确查找】，关键词可以是电话、邮箱、账号
func (l *GetUserByKeywordLogic) GetUserByKeyword(in *luckydraw_pb.GetUserByKeywordReq) (*luckydraw_pb.GetUserByKeywordResp, error) {
	keyword := strings.TrimSpace(in.Keyword)
	if keyword == "" {
		return &luckydraw_pb.GetUserByKeywordResp{}, nil
	}

	mdUser, err := l.svcCtx.UserModel.FindByPassport(l.ctx, keyword)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return &luckydraw_pb.GetUserByKeywordResp{
			User: tools.MdUserToPbUser(mdUser),
		}, nil
	}

	mdUser, err = l.svcCtx.UserModel.FindByEmail(l.ctx, keyword)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return &luckydraw_pb.GetUserByKeywordResp{
			User: tools.MdUserToPbUser(mdUser),
		}, nil
	}

	mdUser, err = l.svcCtx.UserModel.FindByPhone(l.ctx, keyword)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return &luckydraw_pb.GetUserByKeywordResp{
			User: tools.MdUserToPbUser(mdUser),
		}, nil
	}

	return &luckydraw_pb.GetUserByKeywordResp{}, nil
}
