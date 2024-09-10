/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-29 10:22:29
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-07 16:19:09
 * @FilePath: /voting-ranking/api/service/vote.go
 * @Description:
 *
 */
package service

import (
	"voting-ranking/api/dao"
	"voting-ranking/api/dto"
	"voting-ranking/common/result"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IVoteService interface {
	AddVote(c *gin.Context, dto dto.VoteDto)
}

type VoteServiceImpl struct{}

func (v *VoteServiceImpl) AddVote(c *gin.Context, dto dto.VoteDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	user, _ := dao.GetUserByUserId(dto.UserId)
	if user.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "投票用户不存在")
		return
	}

	player, _ := dao.GetPlayerDetail(dto.PlayerId)
	if player.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "参赛选手用户不存在")
		return
	}

	vote, _ := dao.GetVoteInfo(dto)
	if vote.ID != 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "已经投过票了")
		return
	}

	err = dao.AddVote(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "投票失败")
		return
	}
	_, err = dao.UpdatePlayerScore(dto.PlayerId)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "更新选手失败")
		return
	}

	result.Success(c, "投票成功")
}

var voteService = VoteServiceImpl{}

func VoteService() IVoteService {
	return &voteService
}
