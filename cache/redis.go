/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-17 18:26:30
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-17 23:25:40
 * @FilePath: /voting-ranking/cache/redis.go
 * @Description:
 *
 */
package cache

import (
	"context"
	"voting-ranking/config"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDb,
	})
	Rctx = context.Background()
}

func Zscore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}
