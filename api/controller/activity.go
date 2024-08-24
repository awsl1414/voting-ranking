/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-21 23:26:22
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-21 23:30:01
 * @FilePath: /voting-ranking/api/controller/active.go
 * @Description:
 *
 */
package controller

import (
	"voting-ranking/api/dto"
	"voting-ranking/api/service"

	"github.com/gin-gonic/gin"
)

func AddActivity(c *gin.Context) {
	var dto dto.ActivityDto
	c.BindJSON(&dto)
	service.ActivityService().AddActivity(c, dto)
}
