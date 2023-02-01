/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:39:29
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BasePutUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutUserPoolLogic {
	return &BasePutUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BasePutUserPoolLogic) BasePutUserPool(in *luckydraw_pb.PutUserPoolReq) (*luckydraw_pb.PutUserPoolResp, error) {
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
			return &luckydraw_pb.PutUserPoolResp{}, nil
		}
		//如果已经有了
		//那就是违规提交，不予处理
		return nil, errorx.ExternalError(nil, errors.New("该数据可能已经产生变动，请刷新后重试！"))
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

	return &luckydraw_pb.PutUserPoolResp{}, nil
}
