package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserActivitiesLogic {
	return &GetUserActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 某用户可参与的活动
func (l *GetUserActivitiesLogic) GetUserActivities(in *luckydraw_pb.GetUserActivitiesReq) (*luckydraw_pb.GetUserActivitiesResp, error) {
	//策略：先全部查出来，用go语言筛选处理，性能会比mysql语言强一些

	//usersPool里面包含多种pattern的活动
	usersPoolList, err := l.svcCtx.UserPoolModel.FindsByUserId(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	activityIds := []int64{}
	for _, usersPool := range usersPoolList {
		activityIds = append(activityIds, usersPool.ActivityId)
	}

	activityList, err := l.svcCtx.ActivityModel.FindsByIds(l.ctx, activityIds)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//当活动的pattern为所有人时，usersPool中可能会没有录入，需要查出来，加入进去
	//pattern = 1 表示为公开的活动
	activities, err := l.svcCtx.ActivityModel.FindsByPattern(l.ctx, 1)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//开始将两个活动集合进去合并去重处理，把不存在usesPool中的活动筛选出来，赋予默认的usersPool
	var activityMap = make(map[int64]*model.Activity)
	for _, activity := range activityList {
		activityMap[activity.Id] = activity
	}

	//需要注意的是：当pattern为所有人时，积分默认为0，当不入库，只有当客户进行抽奖、或者管理员进行修改之后，才会进行添加入库
	//如果在这里加入usersPool，多个请求时，会产生异步添加冲突
	//所以，只能被动入库，不能主动入库
	//解决方案①：在activity添加端口做修改，当加入一个活动时，主动给所有用户添加usersPool，同时还要在用户注册添加端口，主动加入所有活动的usersPool，缺点：当用户量达到一定量时，添加一个活动，会产生大量的IO工作
	//初衷：解决某下单用户才能在某活动抽奖的问题，这类活动需要将pattern设置为0
	//某用户注册以后，才能看到自己在某活动中的积分情况，所以需要在用户注册添加端口，主动加入所有活动的usersPool
	//如果全部都要入库，会产生大量的数据处理工作，
	//优化方案：只有抽过奖的用户，才会入库usersPool，否则，一律当成默认的处理，
	//总之不能在这里入库

	data := []*luckydraw_pb.UserActivity{}
	for _, usersPool := range usersPoolList {
		for _, activity := range activityList {
			if usersPool.ActivityId == activity.Id {
				data = append(data, &luckydraw_pb.UserActivity{
					Activity: tools.MdActivityToPbActivity(activity),
					UserPool: tools.MdUserPoolToPbUserPool(usersPool),
				})
			}
		}
	}

	var newDefaultUserPool = func(activityId, userId int64) *luckydraw_pb.UserPool {
		return &luckydraw_pb.UserPool{
			Id:         -1,
			Uuid:       uniqueid.NewUUID(),
			ActivityId: activityId,
			UserId:     userId,
			Status:     1,
			Score:      0,
		}
	}

	for _, activity := range activities {
		if _, ok := activityMap[activity.Id]; ok {
			continue
		}
		//不存在usersPool中的活动，赋予默认的usersPool，如果产生修改操作，需在修改端口主动入库
		data = append(data, &luckydraw_pb.UserActivity{
			Activity: tools.MdActivityToPbActivity(activity),
			UserPool: newDefaultUserPool(activity.Id, in.UserId),
		})
	}

	return &luckydraw_pb.GetUserActivitiesResp{
		UserActivities: data,
	}, nil
}
