package adminlogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebByUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWebByUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebByUrlLogic {
	return &GetWebByUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWebByUrlLogic) GetWebByUrl(in *admin_pb.GetWebByUrlReq) (*admin_pb.GetWebByUrlResp, error) {
	weburls := utils.SplitURL(in.Url)
	if len(weburls) == 0 {
		return &admin_pb.GetWebByUrlResp{}, nil
	}

	for _, weburl := range weburls {
		web, err := l.svcCtx.WebModel.FindByUrl(l.ctx, weburl)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		if err == nil {
			return &admin_pb.GetWebByUrlResp{
				Web: dataconv.MdWebToPbWeb(web),
			}, nil
		}
	}

	return &admin_pb.GetWebByUrlResp{}, nil
}
