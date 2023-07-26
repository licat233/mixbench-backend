package ws

import (
	"bytes"
	"context"
	"math"
	"math/rand"
	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/model"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// 预处理消息内容，删除前后空格和换行
func preprocessMessageContent(content string) string {
	return strings.TrimSpace(strings.Trim(content, string(newline)))
}

// 验证载体是否合法
func validatePayload(payload []byte) bool {
	res := bytes.TrimSpace(bytes.ReplaceAll(payload, newline, space))
	return len(res) != 0
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

// 分配客服，核心分配算法
func AssignKefu() int64 {
	kefuCount := OnlineKefu.Len()
	// 前提是有客服在线
	if kefuCount == 0 {
		// 没有客服在线，0表示未分配状态，应该执行消息自动回复机制
		return 0
	}
	OnlineVisitor.RwMutex.RLock()
	// 统计每个 kefuId 对应的 visitorId 的数量
	count := make(map[int64]int64) //kefuId - visitorCount
	for _, kefuId := range OnlineVisitor.Store {
		count[kefuId]++
	}
	OnlineVisitor.RwMutex.RUnlock()

	//可能有kefuId没有分配到onlieVisitor中，所以还需要检查onlineKefu是否有为分配到onlineVisitor中的kefuId
	OnlineKefu.RwMutex.RLock()
	for kefuId := range OnlineKefu.Store {
		if _, ok := count[kefuId]; !ok {
			//如果没有，就统计进去, 0表示没有在线游客分配到该kefu
			count[kefuId] = 0
		}
	}
	OnlineKefu.RwMutex.RUnlock()

	minVisitorCount := int64(math.MaxInt64) // 设置一个较大的初始值
	minKefuId := int64(0)
	for kefuId, visitorCount := range count {
		if visitorCount < minVisitorCount {
			minVisitorCount = visitorCount
			minKefuId = kefuId
		}
	}

	if minKefuId < 1 {
		return 0
	}

	//检查最小游客量的kefu是否有多个，如果有多个，应该采取随机抽取一个
	kefuIds := []int64{}
	for kefuId, visitorCount := range count {
		if visitorCount == minVisitorCount {
			kefuIds = append(kefuIds, kefuId)
		}
	}

	n := len(kefuIds)

	if n == 0 {
		return 0
	} else if n == 1 {
		return kefuIds[0]
	}

	// 创建随机数生成器
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成随机索引并获取随机键
	randomIndex := random.Intn(len(kefuIds))
	randomKefuId := kefuIds[randomIndex]

	return randomKefuId
}
