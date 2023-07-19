/*
 * @Author: licat
 * @Date: 2023-02-20 18:10:59
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:26:00
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"net/http"
	"strconv"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetALuckyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetALuckyListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetALuckyListLogic {
	return &GetALuckyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetALuckyListLogic) GetALuckyList(req *types.GetALuckyListReq) (any, error) {
	//获取游客id
	userIdV := l.r.Header.Get("uid")
	userId, err := strconv.ParseInt(userIdV, 10, 64)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	//活动uuid
	if !govalidator.IsUUID(req.ActivityUUID) {
		return nil, errorx.ErrInvalidUuid
	}
	in := &luckydraw_pb.GetUALuckyDetailsReq{
		UserId:       userId,
		ActivityUuid: req.ActivityUUID,
		ActivityId:   0,
	}
	res, err := l.svcCtx.LuckydrawRpc.LuckydrawClient.GetUALuckyDetails(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.LuckyDetails, err)
}
