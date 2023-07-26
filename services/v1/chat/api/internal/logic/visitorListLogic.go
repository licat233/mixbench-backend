package logic

import (
	"context"
	"errors"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

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
	id, err := jwtx.GetIdByRequest(l.r, "")
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	if id == 0 {
		l.Logger.Error("kefuId == 0")
		return nil, errorx.InternalError(errors.New("current kefuId == 0"))
	}
	//限制在200条
	if req.PageSize > 200 {
		req.PageSize = 200
	}
	//查找出当前客服的好友
	list, total, err := l.svcCtx.VisitorModel.FindsByKefuid(l.ctx, id, req.PageSize, req.Page)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	listData := dataconv.MdVisitors2ApiVisitors(list)
	data := []*types.VisitorListItem{}
	//找出这些在线好友的最后一条消息
	for _, visitor := range listData {
		message, err := l.svcCtx.MessageModel.FindLastByVisitorId(l.ctx, visitor.Id)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Error(err)
			return nil, errorx.InternalError(err)
		}
		item := &types.VisitorListItem{
			LastMessage: dataconv.MdMessageToApiMessage(message),
			Visitor:     visitor,
		}
		data = append(data, item)
	}

	return respx.DefaultListResp(data, total, req.PageSize, req.Page, nil)
}
