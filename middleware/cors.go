/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-09-10 16:32:40
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:32:49
 * @FilePath: /voting-ranking/middleware/cors.go
 * @Description:
 *
 */
// 跨域中间件配置

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRFToken, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-ControlAllow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
