/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-13 17:18:07
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-17 17:34:11
 * @FilePath: /voting-ranking/controllers/vote.go
 * @Description:
 *
 */
package controllers

import (
	"strconv"
	"voting-ranking/models"

	"github.com/gin-gonic/gin"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {

	userId, _ := strconv.Atoi(c.DefaultPostForm("userId", ""))
	playerId, _ := strconv.Atoi(c.DefaultPostForm("playerId", ""))

	if userId == 0 || playerId == 0 {
		ReturnError(c, 4004, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}

	player, _ := models.GetPlayerInfo(playerId)
	if player.Id == 0 {
		ReturnError(c, 4001, "投票选手不存在")
		return
	}

	vote, _ := models.GetVoteInfo(userId, playerId)

	if vote.Id != 0 {
		ReturnError(c, 4001, "已投票")
		return
	}
	rs, err := models.AddVote(userId, playerId)

	if err == nil {
		// 更新参赛选手分数字段，自增一
		models.UpdatePlayerScore(playerId)
		ReturnSuccess(c, 0, "投票成功", rs)
		return
	}
	ReturnError(c, 4004, "请联系管理员")

}
