/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-13 12:06:11
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-18 00:08:34
 * @FilePath: /voting-ranking/controllers/player.go
 * @Description:
 *
 */
package controllers

import (
	"fmt"
	"strconv"
	"time"
	"voting-ranking/cache"
	"voting-ranking/models"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {

	aid, _ := strconv.Atoi(c.DefaultPostForm("aid", "0"))

	rs, err := models.GetPlayers(aid, "id asc")
	fmt.Println(rs)

	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}

	ReturnSuccess(c, 0, "success", rs)

}

func (p PlayerController) GetRanking(c *gin.Context) {
	aid_s := c.DefaultPostForm("aid", "0")

	aid, _ := strconv.Atoi(aid_s)

	redisKey := "ranking:" + aid_s
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		var players []models.Player
		// TODO: 这里应该不能在for循环里做查询
		for _, v := range rs {
			id, _ := strconv.Atoi(v)
			rsInfo, _ := models.GetPlayerInfo(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(c, 0, "r_success", players)
		return

	}

	rsDb, errDb := models.GetPlayers(aid, "score desc")

	if errDb == nil {
		for _, v := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(v.Id, v.Score)).Err()
		}
		// 设置过期时间
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		ReturnSuccess(c, 0, "m_success", rsDb)

		return
	}
	ReturnError(c, 4004, "没有相关信息")
}
