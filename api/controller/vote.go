/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-09-07 16:12:44
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:24:26
 * @FilePath: /voting-ranking/api/controller/vote.go
 * @Description:
 *
 */
package controller

import (
	"voting-ranking/api/dto"
	"voting-ranking/api/service"

	"github.com/gin-gonic/gin"
)

// 投票
// @Summary 投票接口
// @Tags      Vote
// @Produce json
// @Description 投票接口
// @Param data body dto.VoteDto true "data"
// @Success 200 {object} result.Result
// @router /api/vote/add [post]
func AddVote(c *gin.Context) {
	var dto dto.VoteDto
	_ = c.BindJSON(&dto)
	service.VoteService().AddVote(c, dto)
}
