/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 12:25:21
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"net/http"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *TestLogic) Test(req *types.TestReq) (any, error) {
	data := req.Data
	if data == "" {
		data = "test successful"
	}
	resp, err := l.svcCtx.AdminRpc.Test(l.ctx, &admin_pb.TestReq{
		Data: data,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Data, nil)
}
