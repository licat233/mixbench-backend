package logic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VisitorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewVisitorListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *VisitorListLogic {
	return &VisitorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *VisitorListLogic) VisitorList(req *types.VisitorListReq) (any, error) {
	kefuIdV := fmt.Sprintf("%s", l.ctx.Value("KefuId"))
	if kefuIdV == "" {
		l.Logger.Error("not found KefuId in ctx")
		return nil, errorx.InternalError(errors.New("not found KefuId in ctx"))
	}
	kefuId, err := strconv.ParseInt(kefuIdV, 10, 64)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	if kefuId == 0 {
		l.Logger.Error("kefuId == 0")
		return nil, errorx.InternalError(errors.New("current kefuId == 0"))
	}
	list, total, err := l.svcCtx.VisitorModel.FindsByKefuid(l.ctx, req.PageSize, req.Page, kefuId)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	listData := dataconv.MdVisitors2ApiVisitors(list)
	return respx.ListResp("", listData, total, req.PageSize, req.Page), nil
}
