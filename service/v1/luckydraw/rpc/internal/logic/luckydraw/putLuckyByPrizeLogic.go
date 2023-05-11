/*
 * @Author: licat
 * @Date: 2023-02-01 18:51:26
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 19:03:19
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"
	"errors"
	"fmt"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type PutLuckyByPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutLuckyByPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLuckyByPrizeLogic {
	return &PutLuckyByPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过奖品更新幸运单
func (l *PutLuckyByPrizeLogic) PutLuckyByPrize(in *luckydraw_pb.PutLuckyByPrizeReq) (*luckydraw_pb.StatusResp, error) {
	if in.Lucky.Ip != "" && !govalidator.IsIP(in.Lucky.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}
	//先查询出prize_pool_id
	pplist, err := l.svcCtx.PrizePoolModel.FindsByActivityAndPrize(l.ctx, in.ActivityId, in.PrizeId)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	n := len(pplist)
	if n == 0 {
		return nil, errorx.ExtRpcErr(errors.New("the prizesPool is not found "))
	} else if n > 1 {
		e := fmt.Errorf("prizesPool is not unique when activityId=%d,prizeId=%d", in.ActivityId, in.PrizeId)
		l.Logger.Error(e)
		return nil, errorx.IntRpcErr(e)
	}

	lucky := tools.PbLuckyToMdLucky(in.Lucky)
	lucky.PrizePoolId = pplist[0].Id

	if err := l.svcCtx.LuckyModel.Update(l.ctx, lucky); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.StatusResp{}, nil
}
