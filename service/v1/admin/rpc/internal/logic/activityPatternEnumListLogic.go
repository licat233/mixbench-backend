package logic

import (
	"context"

	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPatternEnumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityPatternEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPatternEnumListLogic {
	return &ActivityPatternEnumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityPatternEnumListLogic) ActivityPatternEnumList(in *admin_pb.NilReq) (*admin_pb.EnumListResp, error) {
	mdList := l.svcCtx.ActivityModel.PatternEnumList()
	return &admin_pb.EnumListResp{
		Data: mdEnumListToPbEnumList(mdList),
	}, nil
}
