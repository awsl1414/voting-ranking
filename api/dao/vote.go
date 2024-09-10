/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-29 10:21:55
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-07 16:02:04
 * @FilePath: /voting-ranking/api/dao/vote.go
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

// 投票
func AddVote(dto dto.VoteDto) error {
	vote := model.Vote{
		UserId:     dto.UserId,
		PlayerId:   dto.PlayerId,
		CreateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&vote).Error
	return err

}

// 获取投票信息
func GetVoteInfo(dto dto.VoteDto) (vote model.Vote, err error) {
	err = db.Db.Where("user_id = ? AND player_id = ?", dto.UserId, dto.PlayerId).First(&vote).Error
	return vote, err
}
