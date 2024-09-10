/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-21 15:38:40
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-21 15:38:59
 * @FilePath: /voting-ranking/api/dto/user.go
 * @Description:
 *
 */

package dto

// UserRegisterDto 注册参数
type UserRegisterDto struct {
	Username        string `json:"username" validate:"required"`        // 用户名
	Password        string `json:"password" validate:"required"`        // 密码
	ConfirmPassword string `json:"confirmPassword" validate:"required"` // 密码
}

// UserLoginDto 登录参数
type UserLoginDto struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
}
