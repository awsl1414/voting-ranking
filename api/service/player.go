/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 17:35:51
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:43:57
 * @FilePath: /voting-ranking/api/service/player.go
 * @Description:
 *
 */
package service

import (
	"fmt"
	"voting-ranking/api/dao"
	"voting-ranking/api/dto"
	"voting-ranking/common/result"
	"voting-ranking/common/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var store = util.RedisStore{}

type IPlayerService interface {
	GetPlayerList(c *gin.Context, dto dto.PlayerListDto)
	AddPlayer(c *gin.Context, dto dto.AddPlayerDto)
	GetPlayer(c *gin.Context, id int)
	GetRankList(c *gin.Context, dto dto.PlayerListDto)
}

type PlayerServiceImpl struct{}

// 获取选手列表
func (p *PlayerServiceImpl) GetPlayerList(c *gin.Context, dto dto.PlayerListDto) {
	// 校验参数
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

// 获取选手详情
func (p *PlayerServiceImpl) GetPlayer(c *gin.Context, id int) {
	player, err := dao.GetPlayerDetail(id)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "选手不存在")
		return
	}
	result.Success(c, player)
}

func (p *PlayerServiceImpl) GetRankList(c *gin.Context, dto dto.PlayerListDto) {
	// 参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	// 构建 Redis 缓存的 key，格式为 "rank:aid:<活动ID>"
	redisKey := fmt.Sprintf("rank:aid:%d", dto.Aid)
	// 从缓存中获取排行榜数据
	cacheList := store.Get(redisKey)
	// fmt.Println(cacheList)
	if len(cacheList) > 0 {
		// 直接从缓存中获取玩家详情
		players, err := dao.GetPlayerByIds(cacheList)
		if err != nil {
			result.Failed(c, int(result.ApiCode.FAILED), "获取选手详情失败")
			return
		}
		result.Success(c, players)
		return
	}

	ret, err := dao.GetPlayerList(dto.Aid, "score desc")
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "获取选手列表失败")
		return
	}
	// 将数据库查询到的数据缓存到 Redis 中
	for _, value := range ret {
		store.Set(redisKey, int(value.ID), value.Score)
	}

	result.Success(c, ret)
}

var playerService = PlayerServiceImpl{}

func PlayerService() IPlayerService {
	return &playerService
}
