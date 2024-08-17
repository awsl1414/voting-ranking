/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-13 17:24:30
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-13 21:52:38
 * @FilePath: /voting-ranking/models/vote.go
 * @Description:
 *
 */

package models

import (
	"time"
	"voting-ranking/dao"
)

type Vote struct {
	Id       int   `json:"id"`
	UserId   int   `json:"userId"`
	PlayerId int   `json:"playerId"`
	AddTime  int64 `json:"addTime"`
}

func (Vote) TableName() string {
	return "vote"
}

func GetVoteInfo(userId int, playerId int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id = ? AND player_id = ?", userId, playerId).First(&vote).Error
	return vote, err
}

func AddVote(userId int, playerId int) (int ,error){
	vote := Vote{UserId: userId, PlayerId: playerId, AddTime: time.Now().Unix()}
	err := dao.Db.Create(&vote).Error
	return vote.Id, err
}
