/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 22:58:52
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-12 23:01:24
 * @FilePath: /hnuahe-presentation-voting-ranking/controllers/user.go
 * @Description:
 *
 */
package controllers

import (
	"voting-ranking/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 4001, "两次密码输入不一致")
		return
	}

	user, _ := models.GetUserInfoByUserName(username)

	if user.Id != 0 {
		ReturnError(c, 4001, "用户已存在")
		return
	}
	if _, err := models.AddUser(username, EncryMd5(password)); err != nil {
		ReturnError(c, 4001, "注册失败，请联系管理员")
		return
	}
	ReturnSuccess(c, 0, "注册成功", user)
}
