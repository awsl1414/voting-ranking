/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 23:14:37
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:24:08
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

// 用户注册
// @Summary 用户注册接口
// @Tags      User
// @Produce json
// @Description 用户注册接口
// @Param data body dto.UserRegisterDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/register [post]
func Register(c *gin.Context) {
	var dto dto.UserRegisterDto
	_ = c.BindJSON(&dto)
	service.UserService().Register(c, dto)
}

// 用户登录
// @Summary 用户登录接口
// @Tags      User
// @Produce json
// @Description 用户登录接口
// @Param data body dto.UserLoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/login [post]
func Login(c *gin.Context) {
	var dto dto.UserLoginDto
	_ = c.BindJSON(&dto)
	service.UserService().Login(c, dto)
}
