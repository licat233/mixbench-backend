package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/common/utils"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/mq/types"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LuckyDrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLuckyDrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckyDrawLogic {
	return &LuckyDrawLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 随机抽取奖品
func (l *LuckyDrawLogic) LuckyDraw(in *luckydraw_pb.LuckyDrawReq) (*luckydraw_pb.LuckyDrawResp, error) {
	/**
	 *需要知道哪个用户(user_id)在抽奖，验证该用户状态(status)以及是否已经被删除了(is_deleted)
	 *参加的是哪个活动(activity_id)的抽奖
	 *查找抽奖名单(lucky)，看其抽奖次数是否已经达到活动(activity)的规定上限(max_count)
	 *取出该活动的奖品(prizes)，用于抽奖goodluck
	 *抽取到奖品(prize)以后，建立中奖信息(lucky)
	 *用户(user_id)积分(score)增删操作，如果没有中奖，则给予 10 个积分奖励
	 *最后返回中奖信息(lucky)
	 */

	if !govalidator.IsIP(in.Ip) {
		return nil, errorx.SecRpcErr(errors.New("request ip validator failed"))
	}
	//获取用户信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.SecRpcErr(fmt.Errorf("user not found, id: %d", in.UserId))
	}
	err = model.VerifyUser(user)
	if err != nil {
		return nil, errorx.ExtRpcErr(err)
	}

	//获取活动信息
	mdActivity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.SecRpcErr(fmt.Errorf("activity not found, id: %d", in.ActivityId))
	}
	if mdActivity.Uuid != in.ActivityUuid {
		return nil, errorx.SecRpcErr(errors.New("activity UUID does not match activity ID"))
	}
	err = model.VerifyActivity(mdActivity)
	if err != nil {
		return nil, errorx.ExtRpcErr(err)
	}

	//检查该用户在该活动的抽奖次数
	luckCount, err := l.svcCtx.LuckyModel.FindCountByUserAndActivity(l.ctx, in.UserId, mdActivity.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if luckCount >= mdActivity.MaxCount {
		return nil, errorx.ExtRpcErr(fmt.Errorf("你在[%s]活动中的抽奖次数已达到上限！", mdActivity.Name))
	}

	//获取活动奖品列表
	mdPrizePoolList, err := l.svcCtx.PrizePoolModel.FindByActivityUuid(l.ctx, in.ActivityUuid)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("该活动的奖品池还没有放入奖品"))
	}
	pbPrizePoolList := tools.MdPrizePoolsToPbPrizePools(mdPrizePoolList)

	//抽取奖品
	PrizePool, err := randomPrize(pbPrizePoolList)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//建立中奖信息
	mdLucky := &model.Lucky{
		Uuid:        uniqueid.GenerateUUID(),
		UserId:      user.Id,
		PrizePoolId: PrizePool.Id,
		Status:      0,
		LuckDate:    time.Now(),
		Ip:          in.Ip,
	}
	_, err = l.svcCtx.LuckyModel.Insert(l.ctx, mdLucky)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	/**
	//积分奖励
	// 可以交由mq去处理
	// if PrizePool.IsWin == false {
	// 	user.Score += 10
	// 	if err = l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
	// l.Logger.Error(err)
	// return nil, errorx.IntRpcErr(err)
	// 	}
	// }
	*/

	//查询是否是否中奖
	grade, err := l.svcCtx.PrizeGradeModel.FindOne(l.ctx, PrizePool.GradeId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	isWin := utils.Int64ToBool(grade.IsWin)

	//交给kafka，mq去处理。已经弃用了kafka，使用asynq代替
	// if err := l.kqUpdateUserScore(user.Id, isWin); err != nil {
	// 	l.Logger.Errorf("l.kqUpdateUserScore(user.Id, isWin) error: %v", err)
	// }

	if err := l.asynqUpdateUserScore(user.Id, isWin); err != nil {
		l.Logger.Errorf("消息队列任务处理错误：l.asynqUpdateUserScore(user.Id, isWin) error: %v", err)
	}

	//获取奖品信息
	mdPrize, err := l.svcCtx.PrizeModel.FindOne(l.ctx, PrizePool.PrizeId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//返回中奖信息
	PrizePool.LuckValue = 0
	return &luckydraw_pb.LuckyDrawResp{
		User: tools.MdUserToPbUser(user),
		LuckyDetail: &luckydraw_pb.LuckyDetail{
			Lucky:     tools.MdLuckyToPbLucky(mdLucky),
			PrizePool: PrizePool,
			Activity:  tools.MdActivityToPbActivity(mdActivity),
			Prize:     tools.MdPrizeToPbPrize(mdPrize),
		},
	}, nil
}

func (l *LuckyDrawLogic) asynqUpdateUserScore(userId int64, isWin bool) error {
	if l.svcCtx.AsynqClient == nil {
		err := errors.New("asynq client未初始化")
		return err
	}
	msg := types.LuckdrawResultMsg{
		IsWin:  isWin,
		UserId: userId,
	}
	return l.svcCtx.AsynqClient.AddTask(msg)
}

func (l *LuckyDrawLogic) KqUpdateUserScore(userId int64, isWin bool) error {
	if l.svcCtx.KqueueGoodluckStatusHandlerClient == nil {
		err := errors.New("kafka client 未初始化")
		return err
	}
	m := types.LuckdrawResultMsg{
		IsWin:  isWin,
		UserId: userId,
	}
	body, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}
	return l.svcCtx.KqueueGoodluckStatusHandlerClient.Push(string(body))
}

func randomPrize(pbPrizes []*luckydraw_pb.PrizePool) (*luckydraw_pb.PrizePool, error) {
	//注意：考虑某个傻子设置幸运值为负数，负数或者0的prize不参与抽奖，因为逻辑上完全不可能抽中
	if len(pbPrizes) == 0 {
		return nil, errors.New("randomPrize Error, len(pbPrizes) == 0")
	}
	var totalLuckValue int64
	var prizes []*luckydraw_pb.PrizePool
	for _, prize := range pbPrizes {
		//如果机会 <=0 ，则不参与抽奖
		if prize.LuckValue <= 0 || prize.Total <= 0 {
			continue
		}
		totalLuckValue += prize.LuckValue * prize.Total
		for i := 0; i < int(prize.Total); i++ {
			prizes = append(prizes, prize)
		}
	}

	//如果全都不可能抽中
	if totalLuckValue == 0 {
		return nil, errors.New("randomPrize Error, totalLuckValue = 0")
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Int63n(totalLuckValue)
	temp := random

	sortPrizes(prizes)

	//抽取算法
	for _, prize := range prizes {
		if random < prize.LuckValue {
			return prize, nil
		}
		random -= prize.LuckValue
	}
	return nil, fmt.Errorf("randomPrize Error, totalChance=%d, random=%d", totalLuckValue, temp)
}

// 升序排序
func sortPrizes(pbPrizes []*luckydraw_pb.PrizePool) {
	sort.Slice(pbPrizes, func(i, j int) bool {
		return pbPrizes[i].LuckValue < pbPrizes[j].LuckValue
	})
}
