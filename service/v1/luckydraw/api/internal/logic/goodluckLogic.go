package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"
	"net/http"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/utils"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type GoodluckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGoodluckLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GoodluckLogic {
	return &GoodluckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GoodluckLogic) Goodluck(req *types.GoodluckReq) (any, error) {
	//获取用户id
	userId := ctxdata.GetUidFromCtx(l.ctx)
	//活动uuid检验
	if govalidator.IsUUID(req.ActivityUUID) == false {
		return nil, errorx.ErrInvalidUuid
	}
	ip := utils.RemoteIp(l.r)
	if !govalidator.IsIP(ip) {
		return nil, errorx.ExternalError(errors.New("request ip validator failed"))
	}
	in := &luckydraw_pb.GoodluckReq{
		UserId:       userId,
		ActivityId:   req.ActivityId,
		ActivityUuid: req.ActivityUUID,
		Ip:           ip,
	}
	res, err := l.svcCtx.LuckdrawRpc.Goodluck(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res, err)
}
