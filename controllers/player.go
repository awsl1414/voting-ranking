/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-13 12:06:11
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-13 21:54:07
 * @FilePath: /voting-ranking/controllers/player.go
 * @Description:
 *
 */
package controllers

import (
	"fmt"
	"strconv"
	"voting-ranking/models"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {

	aid, _ := strconv.Atoi(c.DefaultPostForm("aid", ""))

	rs, err := models.GetPlayers(aid)
	fmt.Println(rs)

	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}

	ReturnSuccess(c, 0, "success", rs)

}
