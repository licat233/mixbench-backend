package common

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"mixbench/common/utils"
)

var authIpKeyPrefix = "cache:auth:ip"

func formatCaptchaIpKey(ip string) string {
	return fmt.Sprintf("%s:%v", authIpKeyPrefix, ip)
}

// RequestCountValidator 请求次数验证，每个ip在10分钟内，最多只能maxCount次
func RequestCountValidator(ctx context.Context, redisClient *redis.Redis, r *http.Request, maxCount int) (bool, error) {
	ip := utils.RemoteIp(r)
	key := formatCaptchaIpKey(ip)
	//查询
	countValue, err := redisClient.GetCtx(ctx, key)
	if err != nil {
		return false, err
	}
	if countValue == "" {
		//如果不存在，则创建
		err = redisClient.SetexCtx(ctx, key, "1", 600)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	//如果存在
	count, err := strconv.Atoi(countValue)
	if err != nil {
		//值不符合规范，应该重置
		err = redisClient.SetexCtx(ctx, key, "1", 600)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return count < maxCount, nil
}
