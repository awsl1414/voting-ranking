/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 17:23:19
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-24 17:31:57
 * @FilePath: /voting-ranking/api/model/player.go
 * @Description:
 *
 */
package model

import "voting-ranking/common/util"

// 选手模型对象
type Player struct {
	ID          uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`         // ID
	Aid         int        `gorm:"column:aid;comment:'活动id'" json:"aId"`                         // 活动id
	Ref         string     `gorm:"column:ref;varchar(50);comment:'编号';NOT NULL" json:"ref"`      // 编号
	Nickname    string     `gorm:"column:nickname;varchar(64);comment:'昵称'" json:"nickname"`     // 昵称
	Declaration string     `gorm:"column:declaration;varchar(500);declaration:'宣言'" json:"icon"` // 宣言
	Avatar      string     `gorm:"column:avatar;varchar(64);comment:'头像'" json:"avatar"`         // 头像
	Score       int        `gorm:"column:score;comment:'分数'" json:"score"`                       // 分数
	Phone       string     `gorm:"column:phone;varchar(64);comment:'电话'" json:"phone"`           // 电话
	Note        string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`            // 备注
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
	UpdateTime  util.HTime `gorm:"column:update_time;comment:'创建时间';NOT NULL" json:"updateTime"` // 更新时间
}

// 表名
func (Player) TableName() string {
	return "player"
}
