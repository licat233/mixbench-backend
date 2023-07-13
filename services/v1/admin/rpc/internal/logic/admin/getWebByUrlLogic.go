package adminlogic

import (
	"context"
	"net/url"
	"path"

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
	if in.Url == "" {
		return &admin_pb.GetWebByUrlResp{}, nil
	}
	parsedURL, err := url.Parse(in.Url)
	if err != nil || parsedURL.Host == "" {
		return nil, errorx.ExtRpcErr(err)
	}

	url, _ := url.JoinPath(parsedURL.Host, parsedURL.Path)
	web, err := l.svcCtx.WebModel.FindByUrl(l.ctx, url)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return &admin_pb.GetWebByUrlResp{
			Web: dataconv.MdWebToPbWeb(web),
		}, nil
	}
	//尝试缩短路径查找

	for parentPath := path.Dir(url); parentPath != url; parentPath = path.Dir(url) {
		web, err = l.svcCtx.WebModel.FindByUrl(l.ctx, parentPath)
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

	//不存在，尝试找host
	web, err = l.svcCtx.WebModel.FindByHost(l.ctx, parsedURL.Host)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return &admin_pb.GetWebByUrlResp{
			Web: dataconv.MdWebToPbWeb(web),
		}, nil
	}

	//不存在，尝试查找domain
	web, err = l.svcCtx.WebModel.FindByHost(l.ctx, utils.GetDomain(parsedURL.Host))
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return &admin_pb.GetWebByUrlResp{
			Web: dataconv.MdWebToPbWeb(web),
		}, nil
	}

	return &admin_pb.GetWebByUrlResp{}, nil
}
