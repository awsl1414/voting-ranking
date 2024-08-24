/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 17:35:51
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-24 18:14:48
 * @FilePath: /voting-ranking/api/service/player.go
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

type IPlayerService interface {
	GetPlayerList(c *gin.Context, dto dto.PlayerListDto)
}

type PlayerServiceImpl struct{}

// 获取选手列表
func (p PlayerServiceImpl) GetPlayerList(c *gin.Context, dto dto.PlayerListDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	ret, err := dao.GetPlayerList(dto.Aid, "id asc")
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "选手获取失败")
	}
	result.Success(c, ret)
}

var playerService = PlayerServiceImpl{}

func PlayerService() IPlayerService {
	return &playerService
}
