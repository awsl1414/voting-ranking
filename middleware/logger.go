/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-18 18:31:09
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-18 18:33:39
 * @FilePath: /voting-ranking/middleware/logger.go
 * @Description:
 *
 */
package middleware

import (
	"io"
	"time"
	"voting-ranking/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	logger := log.Log()
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime) / time.Millisecond
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		header := c.Request.Header
		proto := c.Request.Proto
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		err := c.Err()
		body, _ := io.ReadAll(c.Request.Body)
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"header":       header,
			"proto":        proto,
			"err":          err,
			"body":         body,
		}).Info()
	}
}
