/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 17:52:30
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-29 10:16:20
 * @FilePath: /voting-ranking/api/dao/player.go
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

func GetPlayerList(aid int, sort string) ([]model.Player, error) {
	var player []model.Player
	err := db.Db.Where("aid = ?", aid).Order(sort).Find(&player).Error
	return player, err
}

// 添加选手
func AddPlayer(dto dto.AddPlayerDto) (id uint, err error) {
	player := model.Player{
		Aid:         dto.Aid,
		Ref:         dto.Ref,
		Nickname:    dto.Nickname,
		Declaration: dto.Declaration,
		Avatar:      dto.Avatar,
		Score:       dto.Score,
		Phone:       dto.Phone,
		Note:        dto.Note,
		CreateTime:  util.HTime{Time: time.Now()},
		UpdateTime:  util.HTime{Time: time.Now()},
	}
	err = db.Db.Create(&player).Error
	return player.ID, err

}

// 根据电话查询选手
func GetPlayerByPhone(phone string) (player model.Player, err error) {
	err = db.Db.Where("phone = ?", phone).First(&player).Error
	return player, err
}

// 获取选手详情
func GetPlayerDetail(id int) (player model.Player, err error) {
	err = db.Db.Where("id = ?", id).First(&player).Error
	return player, err
}

// 更新选手得分
func UpdatePlayerScore(id int) (player model.Player, err error) {
	err = db.Db.Model(&player).Where("id = ?").UpdateColumn("update_time", util.HTime{Time: time.Now()}).UpdateColumn("score + ?", 1).Error
	return player, err
}
