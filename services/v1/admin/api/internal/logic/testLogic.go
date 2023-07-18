package logic

import (
	"context"
	"fmt"
	"net/http"

	"mixbench/common/respx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"

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
	fmt.Println("你好")
	ip := utils.RemoteIp(l.r)
	addrDetail := utils.QueryAddress(ip)
	return respx.DefaultSingleResp(*addrDetail, nil)
}
