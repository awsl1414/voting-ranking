/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 18:07:43
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-29 10:10:59
 * @FilePath: /voting-ranking/api/controller/player.go
 * @Description:
 *
 */
package controller

import (
	"strconv"
	"voting-ranking/api/dto"
	"voting-ranking/api/service"

	"github.com/gin-gonic/gin"
)

func GetPlayerList(c *gin.Context) {
	var dto dto.PlayerListDto
	c.BindJSON(&dto)
	service.PlayerService().GetPlayerList(c, dto)
}

func AddPlayer(c *gin.Context) {
	var dto dto.AddPlayerDto
	c.BindJSON(&dto)
	service.PlayerService().AddPlayer(c, dto)
}

func GetPlayer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.PlayerService().GetPlayer(c, id)
}
