/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 17:35:51
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-25 23:33:28
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
	AddPlayer(c *gin.Context, dto dto.AddPlayerDto)
}

type PlayerServiceImpl struct{}

// 获取选手列表
func (p *PlayerServiceImpl) GetPlayerList(c *gin.Context, dto dto.PlayerListDto) {
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

// 新增选手
func (p *PlayerServiceImpl) AddPlayer(c *gin.Context, dto dto.AddPlayerDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	// 查询选手是否存在
	player, _ := dao.GetPlayerByPhone(dto.Phone)
	if player.ID > 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "选手已存在")
		return
	}
	_, err = dao.AddPlayer(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "添加失败，请联系管理员")
		return
	}
	result.Success(c, "添加成功")
}

var playerService = PlayerServiceImpl{}

func PlayerService() IPlayerService {
	return &playerService
}
