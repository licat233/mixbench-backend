/*
 * @Author: licat
 * @Date: 2023-02-20 17:40:50
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:25:52
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/utils"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LuckydrawLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewLuckydrawLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LuckydrawLogic {
	return &LuckydrawLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *LuckydrawLogic) Luckydraw(req *types.LuckydrawReq) (any, error) {
	//获取游客id
	userIdV := l.r.Header.Get("uid")
	userId, err := strconv.ParseInt(userIdV, 10, 64)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	//活动uuid检验
	if !govalidator.IsUUID(req.ActivityUUID) {
		return nil, errorx.ErrInvalidUuid
	}
	ip := utils.RemoteIp(l.r)
	if !govalidator.IsIP(ip) {
		return nil, errorx.ExternalError(errors.New("request ip validator failed"))
	}
	in := &luckydraw_pb.LuckydrawReq{
		UserId:       userId,
		ActivityId:   req.ActivityId,
		ActivityUuid: req.ActivityUUID,
		Ip:           ip,
	}
	res, err := l.svcCtx.LuckydrawRpc.Luckydraw(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res, err)
}
