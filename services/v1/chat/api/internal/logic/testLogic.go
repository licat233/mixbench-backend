package logic

import (
	"context"
	"net/http"

	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *TestLogic) Test(req *types.TestReq) (any, error) {
	// ids := []int64{1, 22}
	// data, err := l.svcCtx.VisitorModel.FindsByIds(l.ctx, ids)
	// if err != nil {
	// 	return nil, errorx.InternalError(err)
	// }
	// res := dataconv.MdVisitors2ApiVisitors(data)
	return respx.NewData("hello world"), nil
}
