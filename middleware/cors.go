/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-09-10 16:32:40
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 17:02:18
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
		// 获取请求方法
		method := c.Request.Method
		// 设置允许的跨域请求来源，"*" 表示允许所有域名
		c.Header("Access-Control-Allow-Origin", "*")
		// 设置允许的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRFToken, Authorization, Token")
		// 设置允许的请求方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		// 设置暴露给客户端的头部信息
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-ControlAllow-Origin, Access-Control-Allow-Headers, Content-Type")
		// 允许客户端携带认证信息（如 Cookies）
		c.Header("Access-Control-Allow-Credentials", "true")

		// 对于 OPTIONS 请求方法，直接返回 204 状态码（No Content）
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 继续处理后续请求
		c.Next()
	}
}
