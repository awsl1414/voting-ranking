/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-20 23:05:48
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-25 23:00:38
 * @FilePath: /voting-ranking/api/dao/user.go
 * @Description:
 *
 */
package dao

import (
	"time"
	"voting-ranking/api/dto"
	"voting-ranking/api/model"
	"voting-ranking/common/util"
	"voting-ranking/pkg/db"
)

// 注册
func Register(dto dto.UserRegisterDto) (uint, error) {
	user := model.User{
		Username:   dto.Username,
		Password:   util.EncryptionMd5(dto.Password),
		CreateTime: util.HTime{Time: time.Now()},
		UpdateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&user).Error
	return user.ID, err
}

// 根据用户名查询用户
func GetUserByUserName(username string) (user model.User, err error) {
	err = db.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

// 根据id查询用户
func GetUserByUserId(id int) (user model.User, err error) {
	err = db.Db.Where("id = ?", id).First(&user).Error
	return user, err
}
