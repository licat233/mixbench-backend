package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserEnumsLogic {
	return &GetUserEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户名单枚举列表
func (l *GetUserEnumsLogic) GetUserEnums(in *luckydraw_pb.GetUserEnumsReq) (*luckydraw_pb.Enums, error) {
	list, _, err := l.svcCtx.UserModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	enums := []*luckydraw_pb.Enum{}
	for _, item := range list {
		enums = append(enums, &luckydraw_pb.Enum{
			Label: item.Name,
			Value: item.Id,
		})
	}

	return &luckydraw_pb.Enums{
		Enums: enums,
	}, nil
}
