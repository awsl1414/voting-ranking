/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-09-09 21:23:18
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 15:47:17
 * @FilePath: /voting-ranking/common/util/redisStore.go
 * @Description:
 *
 */
package util

import (
	"context"
	"log"
	"time"
	"voting-ranking/pkg/redis"
)

var ctx = context.Background()

type RedisStore struct{}

// 获取排行榜的值
func (r RedisStore) Get(key string) []string {
	val, err := redis.RedisDb.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil
	}
	return val
}

// 存验证码
func (r RedisStore) Set(key string, id int, score int) error {
	z := redis.ZScore(id, score)
	err := redis.RedisDb.ZAdd(ctx, key, &z).Err()
	if err != nil {
		log.Panicln(err.Error())
		return err
	}
	redis.RedisDb.Expire(ctx, key, time.Hour*1).Err()
	return nil
}
