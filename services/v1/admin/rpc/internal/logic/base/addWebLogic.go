package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddWebLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddWebLogic {
	return &AddWebLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加网站
func (l *AddWebLogic) AddWeb(in *admin_pb.AddWebReq) (*admin_pb.AddWebResp, error) {
	data := &model.Web{
		Uuid:        uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		ProjectId:   in.ProjectId,
		Url:         in.Url,
		Description: in.Description,
		Code:        in.Code,
	}
	if _, err := l.svcCtx.WebModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddWebResp{}, nil
}
