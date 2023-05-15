/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:10:05
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"

	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIdentifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIdentifyLogic {
	return &UserIdentifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 抽奖凭证校验，验证用户身份，颁发jwt
func (l *UserIdentifyLogic) UserIdentify(in *luckydraw_pb.UserIdentifyReq) (*luckydraw_pb.IdentifyResp, error) {
	uLogic := NewGetUserByKeywordLogic(l.ctx, l.svcCtx)

	resp, err := uLogic.GetUserByKeyword(&luckydraw_pb.GetUserByKeywordReq{
		Keyword: in.Passport,
	})
	if err != nil {
		return nil, err
	}
	user := resp.User

	//用户不存在
	if user == nil {
		return &luckydraw_pb.IdentifyResp{
			Status: false,
		}, nil
	}

	//无效用户
	if user.Status != 1 {
		return &luckydraw_pb.IdentifyResp{
			Status: false,
		}, nil
	}

	generate := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	data, err := generate.GenerateToken(&luckydraw_pb.GenerateTokenReq{
		JwtData: &luckydraw_pb.JwtData{
			UserId:   user.Id,
			Ip:       in.JwtData.Ip,
			Platform: in.JwtData.Platform,
		},
	})
	if err != nil {
		return nil, err
	}
	return &luckydraw_pb.IdentifyResp{
		Status:   true,
		JwtToken: data.JwtToken,
	}, nil
}
