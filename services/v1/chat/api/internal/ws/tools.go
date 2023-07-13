package ws

import (
	"context"
	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/model"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// 预处理消息内容
func preprocessMessageContent(content string) string {
	return strings.TrimSpace(strings.ReplaceAll(content, string(newline), string(space)))
}

// TransferVisitor 转移某游客给某客服，返回errorx错误
func TransferVisitor(ctx context.Context, svcCtx *svc.ServiceContext, visitorId, kefuId int64) (err error) {
	logxx := logx.WithContext(ctx)
	if visitorId <= 0 {
		return errorx.ExternalError("visitorId <=0", "visitorId must be greater than 0")
	}

	//修改游客的接待客服

	//判断游客是否存在
	//先判断该游客是否存在
	visitor, err := svcCtx.VisitorModel.FindById(ctx, visitorId)
	if err != nil && err != model.ErrNotFound {
		logxx.Error(err)
		return errorx.InternalError(err)
	}
	//如果不存在，则返回错误
	if err == model.ErrNotFound {
		return errorx.ExternalError("the visitorId does not exist")
	}

	//如果kefuid > 0
	if kefuId > 0 {
		//再判断该客服是否存在
		resp, err := svcCtx.AdminRpc.GetAdminer(ctx, &admin_pb.GetAdminerReq{Id: kefuId})
		if err != nil {
			return err
		}
		//如果不存在，则返回错误
		if resp.Adminer == nil {
			return errorx.ExternalError("the kefuId does not exist")
		}
	}

	//修改该游客的接待客服id
	visitor.KefuId = kefuId
	//更新游客信息
	if err := svcCtx.VisitorModel.Update(ctx, visitor); err != nil {
		logxx.Error(err)
		return errorx.InternalError(err)
	}

	return nil
}
