package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeGradeEnumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeGradeEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeGradeEnumListLogic {
	return &GetPrizeGradeEnumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeGradeEnumListLogic) GetPrizeGradeEnumList(in *admin_pb.EnumListReq) (*admin_pb.EnumListResp, error) {
	list, _, err := l.svcCtx.PrizeGradeModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	data := []*admin_pb.Enum{}
	for _, item := range list {
		data = append(data, &admin_pb.Enum{
			Label: item.Name,
			Value: item.Id,
		})
	}
	return &admin_pb.EnumListResp{
		Data: data,
	}, nil
}
