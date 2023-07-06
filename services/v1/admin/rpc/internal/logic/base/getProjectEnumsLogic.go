package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectEnumsLogic {
	return &GetProjectEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取项目枚举列表
func (l *GetProjectEnumsLogic) GetProjectEnums(in *admin_pb.GetProjectEnumsReq) (*admin_pb.Enums, error) {
	list, _, err := l.svcCtx.ProjectModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	enums := []*admin_pb.Enum{}
	for _, item := range list {
		enums = append(enums, &admin_pb.Enum{
			Label: item.Name,
			Value: item.Id,
		})
	}

	return &admin_pb.Enums{
		Enums: enums,
	}, nil
}
