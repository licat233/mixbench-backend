package ipblack

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpblackListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIpblackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpblackListLogic {
	return &GetIpblackListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIpblackListLogic) GetIpblackList(req *types.GetIpblackListReq) (any, error) {
	in := &model.Ipblack{
		Id: req.Id,
		Ip: req.Ip,
	}
	mdList, total, err := l.svcCtx.IpblackModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Error("failed to query ipblack list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdIpblacks2ApiIpblacks(mdList)

	return respx.DefaultListResp(data, total, req.PageSize, req.Page, nil)
}
