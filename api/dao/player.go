/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 17:52:30
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-24 17:58:05
 * @FilePath: /voting-ranking/api/dao/player.go
 * @Description:
 *
 */
package dao

import (
	"voting-ranking/api/model"
	"voting-ranking/pkg/db"
)

func GetPlayerList(aid int, sort string) ([]model.Player, error) {
	var player []model.Player
	err := db.Db.Where("aid = ?", aid).Order(sort).Find(&player).Error
	return player, err
}
