/*
 * @Author: licat
 * @Date: 2023-02-01 16:43:42
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 17:34:38
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLuckyDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLuckyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLuckyDetailLogic {
	return &GetUserLuckyDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询某用户的中奖信息
func (l *GetUserLuckyDetailLogic) GetUserLuckyDetail(in *luckydraw_pb.GetUserLuckyDetailReq) (*luckydraw_pb.GetUserLuckyDetailResp, error) {
	var user *luckydraw_pb.User
	if in.Passport != "" {
		//查询用户是否存在
		uLogic := NewGetUserByKeywordLogic(l.ctx, l.svcCtx)
		resp, err := uLogic.GetUserByKeyword(&luckydraw_pb.GetUserByKeywordReq{
			Keyword: in.Passport,
		})
		if err != nil {
			return nil, err
		}
		user = resp.User
	} else if in.UserUuid != "" {
		// 查询用户是否存在
		mdUser, err := l.svcCtx.UserModel.FindOneByUuid(l.ctx, in.UserUuid)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		if err == model.ErrNotFound {
			return &luckydraw_pb.GetUserLuckyDetailResp{}, nil
		}
		user = tools.MdUserToPbUser(mdUser)
	} else if in.UserId > 0 {
		// 查询用户是否存在
		mdUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		if err == model.ErrNotFound {
			return &luckydraw_pb.GetUserLuckyDetailResp{}, nil
		}
		user = tools.MdUserToPbUser(mdUser)
	} else {
		return &luckydraw_pb.GetUserLuckyDetailResp{}, nil
	}

	options := model.NewLuckyDetailOptions()
	options.Lucky.UserId = user.Id

	//查询该用户的中奖名单
	mdLuckerDetails, _, err := l.svcCtx.CommonModel.FindsLuckyDetails(l.ctx, options)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	pbLuckerDetails := tools.ConvertLuckyDetails(mdLuckerDetails)
	return &luckydraw_pb.GetUserLuckyDetailResp{
		User:         user,
		LuckyDetails: pbLuckerDetails,
	}, nil
}
