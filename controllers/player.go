package controllers

import (
	"fmt"
	"strconv"
	"voting-ranking/models"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aid_s := c.DefaultPostForm("aid", "")

	aid, _ := strconv.Atoi(aid_s)

	rs, err := models.GetPlayers(aid)
	fmt.Println(rs)

	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", rs)

}
