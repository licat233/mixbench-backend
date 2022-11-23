package logic

import (
	"context"
	"fmt"
	"mixbench/common/utils"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

var CaptchaIpKeyPrefix = "cache:luckdraw:captcha:ip"

func formatCaptchaIpKey(ip string) string {
	return fmt.Sprintf("%s:%v", CaptchaIpKeyPrefix, ip)
}

//RequestCountValid 请求限制，每个ip在10分钟内，最多只能访问1000次
func RequestCountValid(ctx context.Context, redisClient *redis.Redis, r *http.Request) (bool, error) {
	//每个ip在10分钟内，最多只能访问1000次
	ip := utils.RemoteIp(r)
	key := formatCaptchaIpKey(ip)
	//查询是否存在
	exist, err := redisClient.ExistsCtx(ctx, key)
	if err != nil {
		return false, fmt.Errorf("types.go RequestCountValid redisClient.ExistsCtx error: %v", err)
	}
	if exist {
		val, err := redisClient.GetCtx(ctx, key)
		if err != nil {
			return false, fmt.Errorf("types.go RequestCountValid redisClient.GetCtx error: %v", err)
		}
		count, err := strconv.Atoi(val)
		if err != nil {
			return false, fmt.Errorf("types.go RequestCountValid strconv.Atoi error: %v", err)
		}
		if count >= 1000 {
			return false, nil
		}
		count++
		err = redisClient.SetCtx(ctx, key, strconv.Itoa(count))
		if err != nil {
			return false, fmt.Errorf("types.go RequestCountValid redisClient.SetCtx error: %v", err)
		}
	} else {
		err = redisClient.SetexCtx(ctx, key, "1", 10*60)
		if err != nil {
			return false, fmt.Errorf("types.go RequestCountValid redisClient.SetexCtx error: %v", err)
		}
	}

	return true, nil
}
