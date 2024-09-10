/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-18 18:26:30
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:21:17
 * @FilePath: /voting-ranking/router/router.go
 * @Description:
 *
 */
package router

import (
	"voting-ranking/api/controller"
	"voting-ranking/docs"
	"voting-ranking/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	// 创建一个引擎实例
	router := gin.New()

	docs.SwaggerInfo.BasePath = ""

	// 使用 Gin 内置的恢复中间件，捕获所有 panic，并返回500错误（跌机恢复）
	router.Use(gin.Recovery())

	// 日志中间件
	router.Use(middleware.Logger())

	register(router)

	return router

}

func register(router *gin.Engine) {
	// 放行swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}

	player := router.Group("/api/player")
	{
		player.POST("/info", controller.GetPlayer)
		player.POST("/list", controller.GetPlayerList)
		player.POST("/add", controller.AddPlayer)
		player.POST("/rank", controller.GetRankList)

	}
	activity := router.Group("/api/activity")
	{
		activity.POST("/add", controller.AddActivity)
	}

	vote := router.Group("/api/vote")
	{
		vote.POST("/add", controller.AddVote)
	}
}
