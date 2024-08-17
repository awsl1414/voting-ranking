/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-13 12:09:03
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-13 21:57:31
 * @FilePath: /voting-ranking/models/player.go
 * @Description:
 *
 */
package models

import (
	"voting-ranking/dao"

	"gorm.io/gorm"
)

type Player struct {
	Id          int    `json:"id"`
	Aid         int    `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nickname"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       string `json:"score"`
}

func (Player) TableName() string {
	return "player"
}

func GetPlayerInfo(id int) (Player, error) {
	var player Player
	err := dao.Db.Where("id = ?", id).First(&player).Error
	return player, err
}

func GetPlayers(aid int) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid = ?", aid).Find(&players).Error
	return players, err
}

func UpdatePlayerScore(id int) {
	var player Player
	dao.Db.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}
