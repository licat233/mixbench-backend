package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/luckdraw/rpc/internal/svc"
	"mixbench/service/v1/luckdraw/rpc/luckdraw_pb"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"
	"mixbench/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectPrizesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPrizesLogic {
	return &SelectPrizesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/** SelectPrizes 用户选择要领取的奖品
 */
func (l *SelectPrizesLogic) SelectPrizes(in *luckdraw_pb.SelectPrizesReq) (*luckdraw_pb.SelectPrizesResp, error) {
	if len(in.LuckerUuids) == 0 {
		return nil, errorx.SecRpcErr(errors.New("luckerUuidList len = 0"))
	}
	//查询是否存在该活动
	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.SecRpcErr(errors.New("activity not found"))
	}
	if activity.Uuid != in.ActivityUuid {
		return nil, errorx.SecRpcErr(errors.New("activity UUID does not match activity ID"))
	}

	//取出中奖名单
	luckerList, err := l.svcCtx.LuckerModel.FindByUserAndActivity(l.ctx, in.UserId, activity.Id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.SecRpcErr(errors.New("lucker not found"))
	}
	//uuids去重
	inUuidList := utils.RemoveDuplicatesForString(in.LuckerUuids)

	//已领取的奖品
	var selectedLuckerList []*model.Lucker   //已领取的奖品
	var noSelectedLuckerList []*model.Lucker //未领取的奖品
	var mdUuidList []string
	var selectedUuidList []string
	for _, lucker := range luckerList {
		if lucker.Status == 1 {
			selectedLuckerList = append(selectedLuckerList, lucker)
			mdUuidList = append(mdUuidList, lucker.Uuid)
			selectedUuidList = append(selectedUuidList, lucker.Uuid)
		} else if lucker.Status == 0 {
			noSelectedLuckerList = append(noSelectedLuckerList, lucker)
			mdUuidList = append(mdUuidList, lucker.Uuid)
		}
	}

	if len(noSelectedLuckerList) == 0 {
		return nil, errorx.SecRpcErr(errors.New("the user has no prizes to claim"))
	}

	if len(selectedLuckerList) >= int(activity.MaxPrizes) {
		return nil, errorx.SecRpcErr(errors.New("the user's prizes in the current activity have been collected"))
	}

	//校验是否有多余的uuid传进来
	if !utils.SliceContainString(mdUuidList, inUuidList...) {
		return nil, errorx.SecRpcErr(errors.New("luckerUuidList存在多余的uuid"))
	}

	//此时inUuidList内可能含有已经领取过的uuid，也有可能全部是已经领取过得uuid
	//合并
	inUuidList = append(inUuidList, selectedUuidList...)
	//去重
	inUuidList = utils.RemoveDuplicatesForString(inUuidList)
	//判断是否会超额领取
	if len(inUuidList) > int(activity.MaxPrizes) {
		return nil, errorx.SecRpcErr(errors.New("the number of prizes received exceeds the activity limit"))
	}

	//过滤已经未领取的奖品
	uuidList := utils.SliceFilterString(inUuidList, selectedUuidList...)
	if len(uuidList) == 0 {
		return nil, errorx.SecRpcErr(errors.New("there is no prize to claim"))
	}

	//更新领取状态
	err = l.svcCtx.LuckerModel.UpdateSomeStatusByUuid(l.ctx, uuidList, 1)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//如果领取的奖品已达到限额值，其它奖品则为弃领状态
	//取出中奖名单
	newestLuckerList, err := l.svcCtx.LuckerModel.FindByUserAndActivity(l.ctx, in.UserId, activity.Id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.SecRpcErr(errors.New("there is no prize to claim"))
	}

	var selectNum int64
	var noIds []int64
	for _, lucker := range newestLuckerList {
		if lucker.Status == 1 {
			selectNum++
		} else if lucker.Status == 0 {
			noIds = append(noIds, lucker.Id)
		}
	}
	//如果已经领取完毕，其它奖品就等于弃领了
	if selectNum >= activity.MaxPrizes {
		err = l.svcCtx.LuckerModel.UpdateSomeStatusById(l.ctx, noIds, -1)
		if err != nil {
			logx.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
	}

	//查询用户已领取的中奖产品
	mdPrizeList, err := l.svcCtx.PrizeModel.FindsByUserAndActivityAndStatus(l.ctx, in.UserId, activity.Id, 1)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	pbPrizes := convertPrizes(mdPrizeList)
	return &luckdraw_pb.SelectPrizesResp{
		Prizes: pbPrizes,
	}, nil
}
