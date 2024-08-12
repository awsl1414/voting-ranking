/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 22:01:28
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-12 23:13:30
 * @FilePath: /hnuahe-presentation-voting-ranking/router/router.go
 * @Description:
 *
 */
package router

import (
	"voting-ranking/controllers"
	"voting-ranking/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	test := r.Group("/test")

	test.GET("", controllers.Test{}.Hello)
	test.GET("/log", controllers.Test{}.LogTest)

	user := r.Group("/user")
	user.POST("/register", controllers.UserController{}.Register)
	user.POST("/login", controllers.UserController{}.Login)

	return r

}
