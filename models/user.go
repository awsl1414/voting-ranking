/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 23:04:34
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-12 22:57:23
 * @FilePath: /hnuahe-presentation-voting-ranking/models/user.go
 * @Description:
 *
 */
package models

import (
	"time"
	"voting-ranking/dao"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"addTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUserName(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func AddUser(username string, password string) (int, error) {
	user := User{Username: username, Password: password, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
