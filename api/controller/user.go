/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 23:14:37
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-21 22:14:20
 * @FilePath: /voting-ranking/api/controller/user.go
 * @Description:
 *
 */

package controller

import (
	"voting-ranking/api/dto"
	"voting-ranking/api/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var dto dto.UserRegisterDto
	_ = c.BindJSON(&dto)
	service.UserService().Register(c, dto)
}

func Login(c *gin.Context) {
	var dto dto.UserLoginDto
	_ = c.BindJSON(&dto)
	service.UserService().Login(c, dto)
}
