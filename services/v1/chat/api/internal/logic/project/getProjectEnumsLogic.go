package project

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectEnumsLogic {
	return &GetProjectEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectEnumsLogic) GetProjectEnums(req *types.GetProjectEnumsReq) (any, error) {
	mdProjectList, err := l.svcCtx.ProjectModel.FindAll(l.ctx)
	if err != nil {
		l.Logger.Error("failed to query project list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := make([]*types.Enum, 0)
	for _, v := range mdProjectList {
		data = append(data, &types.Enum{
			Label: v.Name,
			Value: v.Id,
		})
	}

	return respx.DefaultSingleResp(data, nil)
}
