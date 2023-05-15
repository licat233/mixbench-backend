package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityEnumsLogic {
	return &GetActivityEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动枚举列表
func (l *GetActivityEnumsLogic) GetActivityEnums(in *luckydraw_pb.GetActivityEnumsReq) (*luckydraw_pb.Enums, error) {
	list, _, err := l.svcCtx.ActivityModel.FindList(l.ctx, -1, -1, "", nil)
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
