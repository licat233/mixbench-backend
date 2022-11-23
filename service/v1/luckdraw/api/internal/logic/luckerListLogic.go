package logic

import (
	"context"
	"mixbench/service/v1/luckdraw/api/internal/svc"
	"mixbench/service/v1/luckdraw/api/internal/types"
	"mixbench/service/v1/luckdraw/rpc/luckdraw_pb"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LuckerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckerListLogic {
	return &LuckerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//获取自己在某活动的中奖名单
func (l *LuckerListLogic) LuckerList(req *types.LuckerListReq) (resp any, err error) {
	//获取用户id
	userId := ctxdata.GetUidFromCtx(l.ctx)
	//活动uuid
	if govalidator.IsUUID(req.ActivityUUID) == false {
		return nil, errorx.ErrInvalidUuid
	}
	in := &luckdraw_pb.GetLuckerListReq{
		UserId:       userId,
		ActivityUuid: req.ActivityUUID,
	}
	res, err := l.svcCtx.LuckdrawRpc.GetLuckerList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.LuckerDetails, err)
}
