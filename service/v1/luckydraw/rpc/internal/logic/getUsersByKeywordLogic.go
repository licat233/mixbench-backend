/*
 * @Author: licat
 * @Date: 2023-02-01 17:00:32
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 17:19:28
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersByKeywordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersByKeywordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersByKeywordLogic {
	return &GetUsersByKeywordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过关键词查找用户【模糊查找】，关键词可以是电话、邮箱、账号、姓名等等
func (l *GetUsersByKeywordLogic) GetUsersByKeyword(in *luckydraw_pb.GetUsersByKeywordReq) (*luckydraw_pb.GetUsersByKeywordResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	if in.ListReq.Keyword == "" {
		return &luckydraw_pb.GetUsersByKeywordResp{}, nil
	}

	list, total, err := l.svcCtx.UserModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, nil)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetUsersByKeywordResp{
		Users: tools.MdUsersToPbUsers(list),
		Total: total,
	}, nil
}
