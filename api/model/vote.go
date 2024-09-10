/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-29 10:18:07
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-29 10:20:35
 * @FilePath: /voting-ranking/api/model/vote.go
 * @Description:
 *
 */
package model

import "voting-ranking/common/util"

// 活动投票模型对象
type Vote struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`         //ID
	UserId     int        `gorm:"column:user_id;comment:'用户账号';NOT NULL" json:"userId"`         // 用户账号
	PlayerId   int        `gorm:"column:player_id;comment:'密码';NOT NULL" json:"playId"`         // 密码
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
}

func (Vote) TableName() string {
	return "vote"
}
