package visitor

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVisitorEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVisitorEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVisitorEnumsLogic {
	return &GetVisitorEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVisitorEnumsLogic) GetVisitorEnums(req *types.GetVisitorEnumsReq) (any, error) {
	mdVisitorList, err := l.svcCtx.VisitorModel.FindAll(l.ctx)
	if err != nil {
		l.Logger.Error("failed to query visitor list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := make([]*types.Enum, 0)
	for _, v := range mdVisitorList {
		data = append(data, &types.Enum{
			Label: v.Username,
			Value: v.Id,
		})
	}

	return respx.DefaultSingleResp(data, nil)
}
