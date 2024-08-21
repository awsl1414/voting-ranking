/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 23:16:50
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-19 23:30:42
 * @FilePath: /voting-ranking/api/model/user.go
 * @Description:
 *
 */

package model

import "voting-ranking/common/util"

// 用户模型对象
type User struct {
	ID       uint   `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                //ID
	Username string `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"` // 用户账号
	Password string `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`   // 密码

	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
	UpdateTime util.HTime `gorm:"column:update_time;comment:'创建时间';NOT NULL" json:"updateTime"` // 更新时间
}

// 表名
func (User) TableName() string {
	return "user"
}
