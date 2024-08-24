/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-18 18:26:30
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-24 18:16:09
 * @FilePath: /voting-ranking/router/router.go
 * @Description:
 *
 */
package router

import (
	"voting-ranking/api/controller"
	"voting-ranking/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 创建一个引擎实例
	router := gin.New()
	// 使用 Gin 内置的恢复中间件，捕获所有 panic，并返回500错误（跌机恢复）
	router.Use(gin.Recovery())

	// 日志中间件
	router.Use(middleware.Logger())

	register(router)

	return router

}

func register(router *gin.Engine) {

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}

	player := router.Group("/api/player")
	{
		player.POST("/list", controller.GetPlayerList)
	}
	activity := router.Group("/api/activity")
	{
		activity.POST("/add", controller.AddActivity)
	}
}
