package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/logic/common"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizeGradeEnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizeGradeEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizeGradeEnumListLogic {
	return &PrizeGradeEnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizeGradeEnumListLogic) PrizeGradeEnumList() (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizeGradeEnumList(l.ctx, &admin_pb.EnumListReq{})
	if err != nil {
		return nil, err
	}
	data := common.ConvertEnumList(resp.Data)
	return respx.DefaultSingleResp(data, err)
	// return common.ConvertEnumList(resp.Data), nil
}
