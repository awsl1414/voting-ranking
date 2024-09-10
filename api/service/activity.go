/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-21 23:14:56
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-29 10:25:31
 * @FilePath: /voting-ranking/api/service/activity.go
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

type IActivityService interface {
	AddActivity(c *gin.Context, dto dto.ActivityDto)
}

type ActivityServiceImpl struct{}

func (a *ActivityServiceImpl) AddActivity(c *gin.Context, dto dto.ActivityDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	activity := dao.GetActivityByName(dto.Name)

	if activity.ID > 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "活动已存在")
		return
	}

	err = dao.AddActivity(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "活动创建失败")
		return
	}
	result.Success(c, "活动创建成功")
}

var activityService = ActivityServiceImpl{}

func ActivityService() IActivityService {
	return &activityService
}
