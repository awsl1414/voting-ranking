/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-21 22:49:42
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-21 22:57:10
 * @FilePath: /voting-ranking/api/model/active.go
 * @Description:
 *
 */
package model

import "voting-ranking/common/util"

// 活动模型
type Activity struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`         // ID
	Name       string     `gorm:"column:name;varchar(64);comment:'活动名称';NOT NULL" json:"name"`  // 活动名称
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 添加时间
}

func (Activity) TableName() string {
	return "activity"
}
