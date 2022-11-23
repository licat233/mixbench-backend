package logic

import (
	"context"
	"errors"
	"fmt"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type PutLuckerByPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutLuckerByPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLuckerByPrizeLogic {
	return &PutLuckerByPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutLuckerByPrizeLogic) PutLuckerByPrize(in *admin_pb.LuckerPutByPrizeReq) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}
	if in.Ip != "" && !govalidator.IsIP(in.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}
	//先查询出prizes_pool_id
	pplist, err := l.svcCtx.PrizesPoolModel.FindsByActivityAndPrize(l.ctx, in.ActivityId, in.PrizeId)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	n := len(pplist)
	if n == 0 {
		return nil, errorx.ExtRpcErr(errors.New("not found prizesPool"))
	} else if n > 1 {
		e := fmt.Errorf("prizesPool is not unique when activityId=%d,prizeId=%d", in.ActivityId, in.PrizeId)
		logx.Error(e)
		return nil, errorx.IntRpcErr(e)
	}

	lukeDate := utils.TimestampHandle(in.LuckDate)
	data := &model.Lucker{
		Id:           in.Id,
		Uuid:         in.Uuid,
		UserId:       in.UserId,
		PrizesPoolId: pplist[0].Id,
		Status:       in.Status,
		LuckDate:     utils.Int64ToTime(lukeDate),
		Ip:           in.Ip,
	}

	if err := l.svcCtx.LuckerModel.Update(l.ctx, data); err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
