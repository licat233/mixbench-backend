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

type PutUsersPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutUsersPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUsersPoolLogic {
	return &PutUsersPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutUsersPoolLogic) PutUsersPool(in *admin_pb.UsersPool) (*admin_pb.StatusResp, error) {
	//约定好的id=-1就表示是没有入库的usersPool
	if in.Id == -1 {
		//防止重复添加
		_, err := l.svcCtx.UsersPoolModel.FindByUidAndAid(l.ctx, in.UserId, in.ActivityId)
		if err != nil && err != model.ErrNotFound {
			logx.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		//没有这样的用户池
		if err == model.ErrNotFound {
			//主动添加上去
			if _, e := l.svcCtx.UsersPoolModel.Insert(l.ctx, &model.UsersPool{
				Uuid:       uniqueid.NewUUID(),
				ActivityId: in.ActivityId,
				UserId:     in.UserId,
				Status:     in.Status,
				Score:      in.Score,
			}); e != nil {
				logx.Error(e)
				return nil, errorx.InternalError(e)
			}
			return newStatusResponse(), nil
		}
		//如果已经有了
		//那就是违规提交，不予处理
		return nil, errorx.ExternalError(nil, errors.New("该可能已经已经产生变动，请刷新后重新修改！"))
	}

	data := &model.UsersPool{
		Id:         in.Id,
		Uuid:       in.Uuid,
		ActivityId: in.ActivityId,
		UserId:     in.UserId,
		Status:     in.Status,
		Score:      in.Score,
	}

	if err := l.svcCtx.UsersPoolModel.Update(l.ctx, data); err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
