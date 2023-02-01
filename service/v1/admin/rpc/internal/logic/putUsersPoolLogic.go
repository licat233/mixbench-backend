package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUserPoolLogic {
	return &PutUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutUserPoolLogic) PutUserPool(in *admin_pb.UserPool) (*admin_pb.StatusResp, error) {
	//约定好的id=-1就表示是没有入库的usersPool
	if in.Id == -1 {
		//防止重复添加
		_, err := l.svcCtx.UserPoolModel.FindByUidAndAid(l.ctx, in.UserId, in.ActivityId)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		//没有这样的用户池
		if err == model.ErrNotFound {
			//主动添加上去
			if _, e := l.svcCtx.UserPoolModel.Insert(l.ctx, &model.UserPool{
				Uuid:       uniqueid.NewUUID(),
				ActivityId: in.ActivityId,
				UserId:     in.UserId,
				Status:     in.Status,
				Score:      in.Score,
			}); e != nil {
				l.Logger.Error(e)
				return nil, errorx.InternalError(e)
			}
			return newStatusResponse(), nil
		}
		//如果已经有了
		//那就是违规提交，不予处理
		return nil, errorx.ExternalError(nil, errors.New("该可能已经已经产生变动，请刷新后重新修改！"))
	}

	data := &model.UserPool{
		Id:         in.Id,
		Uuid:       in.Uuid,
		ActivityId: in.ActivityId,
		UserId:     in.UserId,
		Status:     in.Status,
		Score:      in.Score,
	}

	if err := l.svcCtx.UserPoolModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
