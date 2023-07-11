package redisx

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

// KeyValueValidator 验证key值在某一时间内的有效性
// expire 单位second
// 当 err == nil 时，表示验证通过
func KeyValueValidator(ctx context.Context, redisClient *redis.Redis, key string, expire int, maxCount int) (isExtErr bool, err error) {
	if key = strings.TrimSpace(key); len(key) == 0 {
		return true, errors.New("redis key cannot be equal to nil")
	}
	if maxCount <= 0 {
		return false, errors.New("maxCount cannot <= 0")
	}
	if expire <= 0 {
		return false, errors.New("expire cannot <= 0")
	}
	// 查询
	countValue, err := redisClient.GetCtx(ctx, key)
	if err != nil {
		return false, err
	}
	//如果不存在
	if countValue == "" {
		//创建
		if err = redisClient.SetexCtx(ctx, key, "1", expire); err != nil {
			return false, err
		}
		//第一次访问。放行
		return false, nil
	}
	//如果存在
	count, err := strconv.Atoi(countValue)
	if err != nil {
		//值不符合规范，应该重置，有效期1分钟
		if err = redisClient.SetexCtx(ctx, key, "1", expire); err != nil {
			return false, err
		}
		//已重置，放行
		return false, nil
	}

	if count > maxCount {
		ttl, err := redisClient.TtlCtx(ctx, key)
		if err != nil {
			return false, err
		}
		if ttl > 0 {
			return true, fmt.Errorf("too frequent operations, please try again in %d seconds", ttl)
		}
		// ttl如果刚好等于0，表示已经可以再次访问了
		return false, nil
	}

	count++
	countValue = strconv.Itoa(count)
	if err = redisClient.SetCtx(ctx, key, countValue); err != nil {
		return false, err
	}

	return false, nil
}
