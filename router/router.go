/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 22:01:28
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-11 22:20:42
 * @FilePath: /hnuahe-presentation-voting-ranking/router/router.go
 * @Description:
 *
 */
package router

import (
	"voting-ranking/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	test := r.Group("/test")

	test.GET("", controllers.Test{}.Hello)

	return r

}
