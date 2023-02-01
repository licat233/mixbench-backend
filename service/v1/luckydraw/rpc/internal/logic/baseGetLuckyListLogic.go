/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:37:58
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetLuckyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetLuckyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetLuckyListLogic {
	return &BaseGetLuckyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetLuckyListLogic) BaseGetLuckyList(in *luckydraw_pb.GetLuckyListReq) (*luckydraw_pb.GetLuckyListResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	if in.Lucky != nil && in.Lucky.Ip != "" && !govalidator.IsIP(in.Lucky.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}

	list, total, err := l.svcCtx.LuckyModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, tools.PbLuckyToMdLucky(in.Lucky))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.GetLuckyListResp{
		Luckies: tools.MdLuckyListToPbLuckyList(list),
		Total:   total,
	}, nil
}
