/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 11:21:05
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/logic/dataconv"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *CurrentUserLogic {
	return &CurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *CurrentUserLogic) CurrentUser() (any, error) {
	id, err := jwtx.GetIdByRequest(l.r, "")
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	rpcResp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.GetAdminerReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	if rpcResp.Adminer == nil {
		return nil, errorx.New("user does not exist")
	}
	pbAdminer := rpcResp.Adminer
	data := dataconv.PbAdminerToApiAdminer(pbAdminer)
	if data != nil {
		data.Password = ""
	}

	return respx.DefaultSingleResp(data, nil)
}
