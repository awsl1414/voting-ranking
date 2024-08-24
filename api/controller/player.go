package controller

import (
	"voting-ranking/api/dto"
	"voting-ranking/api/service"

	"github.com/gin-gonic/gin"
)

func GetPlayerList(c *gin.Context) {
	var dto dto.PlayerListDto
	c.BindJSON(&dto)
	service.PlayerService().GetPlayerList(c, dto)
}
