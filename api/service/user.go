/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 23:15:56
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-25 23:31:52
 * @FilePath: /voting-ranking/api/service/user.go
 * @Description: 用户服务层
 *
 */
// 用户服务层
package service

import (
	"voting-ranking/api/dao"
	"voting-ranking/api/dto"
	"voting-ranking/common/result"
	"voting-ranking/common/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 用户服务接口
type IUserService interface {
	Register(c *gin.Context, dto dto.UserRegisterDto)
	Login(c *gin.Context, dto dto.UserLoginDto)
}

type UserServiceImpl struct{}

// 注册
func (u UserServiceImpl) Register(c *gin.Context, dto dto.UserRegisterDto) {

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
	user, _ := dao.GetUserByUserName(dto.Username)
	if user.ID > 0 {
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

// 登陆
func (u UserServiceImpl) Login(c *gin.Context, dto dto.UserLoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	user, _ := dao.GetUserByUserName(dto.Username)

	if user.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户或密码不正确")
		return
	}

	if user.Password != util.EncryptonMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.FAILED), "用户或密码不正确")
		return
	}
	result.Success(c, user)
}

var userService = UserServiceImpl{}

// 面向接口编程
func UserService() IUserService {
	return &userService
}
