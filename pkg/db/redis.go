/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-09-09 18:34:51
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-09 21:12:31
 * @FilePath: /voting-ranking/pkg/db/redis.go
 * @Description:
 *
 */
package db

import (
	"context"
	"voting-ranking/common/config"

	"github.com/go-redis/redis/v8"
)

var RedisDb *redis.Client

// SetupRedis 连接到 Redis 数据库
func SetupRedis() error {
	var ctx = context.Background() // 创建上下文，用于 Redis 操作

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       1,
	})
	// 检查 Redis 服务器是否可用
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// ZScore 创建一个 Redis 的 ZSet 元素
func ZScore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}
