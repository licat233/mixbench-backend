/*
 * @Author: licat
 * @Date: 2023-02-20 18:10:59
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:11:20
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/common/utils"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIdentifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUserIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UserIdentifyLogic {
	return &UserIdentifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UserIdentifyLogic) UserIdentify(req *types.UserIdentifyReq) (any, error) {
	ip := utils.RemoteIp(l.r)
	platform, err := utils.RemotePlatform(l.r)
	if err != nil {
		return nil, errorx.ErrAuth
	}
	in := &luckydraw_pb.UserIdentifyReq{
		Passport: req.Passport,
	}
	resp, err := l.svcCtx.LuckydrawRpc.GetUserByKeyword(l.ctx, &luckydraw_pb.GetUserByKeywordReq{
		Keyword: in.Passport,
	})
	if err != nil {
		return nil, err
	}
	if resp.User == nil {
		return nil, errorx.ErrAuth
	}
	id := resp.User.Id
	claims := jwtx.NewClaims()
	claims[jwtx.JwtUidKey] = id
	claims[jwtx.JwtIpKey] = ip
	claims[jwtx.JwtPlatformKey] = platform
	token, err := jwtx.CreateToken(claims, l.svcCtx.Config.VisitorAuthSecret)
	if err != nil {
		return nil, errorx.InternalError(err)
	}

	data := &types.JwtToken{
		AccessToken:  token,
		AccessExpire: 0,
		RefreshAfter: 0,
	}

	return respx.DefaultSingleResp(data, err)
}
