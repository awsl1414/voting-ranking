/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-09-07 16:12:44
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-07 16:14:01
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

func AddVote(c *gin.Context) {
	var dto dto.VoteDto
	_ = c.BindJSON(&dto)
	service.VoteService().AddVote(c, dto)
}
