/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 18:07:43
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:28:30
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

// 获取选手列表
// @Tags      Player
// @Summary 获取选手列表接口
// @Produce json
// @Description 获取选手列表接口
// @Param data body dto.PlayerListDto true "data"
// @Success 200 {object} result.Result
// @router /api/player/list [post]
func GetPlayerList(c *gin.Context) {
	var dto dto.PlayerListDto
	c.BindJSON(&dto)
	service.PlayerService().GetPlayerList(c, dto)
}

// 新增选手信息
// @Summary 更新选手信息接口
// @Tags      Player
// @Produce json
// @Description 更新选手信息接口
// @Param data body dto.AddPlayerDto true "data"
// @Success 200 {object} result.Result
// @router /api/player/add [post]
func AddPlayer(c *gin.Context) {
	var dto dto.AddPlayerDto
	c.BindJSON(&dto)
	service.PlayerService().AddPlayer(c, dto)
}

// 获取选手详情
// @Summary 获取选手详情接口
// @Tags      Player
// @Produce json
// @Description 获取选手详情接口
// @Param id query int true "Id"
// @Success 200 {object} result.Result
// @router /api/player/info [post]
func GetPlayer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.PlayerService().GetPlayer(c, id)
}

// 获取排行傍
// @Tags      Player
// @Summary 获取排行傍列表接口
// @Produce json
// @Description 获取排行傍列表接口
// @Param data body dto.PlayerListDto true "data"
// @Success 200 {object} result.Result
// @router /api/player/rank [post]
func GetRankList(c *gin.Context) {
	var dto dto.PlayerListDto
	_ = c.BindJSON(&dto)
	service.PlayerService().GetRankList(c, dto)
}
