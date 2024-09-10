/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-21 23:26:22
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:09:22
 * @FilePath: /voting-ranking/api/controller/activity.go
 * @Description:
 *
 */
package controller

import (
	"voting-ranking/api/dto"
	"voting-ranking/api/service"

	"github.com/gin-gonic/gin"
)

// 创建活动
// @Summary 创建活动接口
// @Tags      Activity
// @Produce json
// @Description 创建活动接口
// @Param data body dto.ActivityDto true "data"
// @Success 200 {object} result.Result
// @router /api/activity/add [post]
func AddActivity(c *gin.Context) {
	var dto dto.ActivityDto
	c.BindJSON(&dto)
	service.ActivityService().AddActivity(c, dto)
}
