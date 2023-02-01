package logic

import (
	"context"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LuckyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckyListLogic {
	return &LuckyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取自己在某活动的中奖名单
func (l *LuckyListLogic) LuckyList(req *types.LuckyListReq) (resp any, err error) {
	//获取用户id
	userId := ctxdata.GetUidFromCtx(l.ctx)
	//活动uuid
	if govalidator.IsUUID(req.ActivityUUID) == false {
		return nil, errorx.ErrInvalidUuid
	}
	in := &luckydraw_pb.GetLuckyListReq{
		UserId:       userId,
		ActivityUuid: req.ActivityUUID,
	}
	res, err := l.svcCtx.LuckdrawRpc.GetLuckyList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.LuckyDetails, err)
}
