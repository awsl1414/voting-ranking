/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 23:15:56
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-21 16:34:45
 * @FilePath: /voting-ranking/api/service/user.go
 * @Description: 用户服务层
 *
 */
// 用户服务层
package service

import (
	"fmt"
	"voting-ranking/api/dao"
	"voting-ranking/api/dto"
	"voting-ranking/common/result"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 用户服务接口
type IUserService interface {
	Register(c *gin.Context, dto dto.UserRegisterDto)
}

type UserServiceImpl struct{}

// 注册
func (u UserServiceImpl) Register(c *gin.Context, dto dto.UserRegisterDto) {
	fmt.Println(dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	if dto.Password != dto.ConfirmPassword {
		result.Failed(c, int(result.ApiCode.FAILED), "两次密码不一致")
		return
	}
	// 判断用户是否存在
	userByName := dao.GetUserByUserName(dto.Username)
	if userByName.ID > 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户已存在")
		return
	}
	_, err = dao.Register(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "注册失败，请联系管理员")
		return
	}
	result.Success(c, "注册成功")
}

var userService = UserServiceImpl{}

// 面向接口编程
func UserService() IUserService {
	return &userService
}
