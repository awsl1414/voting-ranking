package controllers

import (
	"voting-ranking/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Test struct{}

func (t Test) Hello(c *gin.Context) {
	ReturnSuccess(c, 200, "请求成功", "hello")
}

func (t Test) LogTest(c *gin.Context) {
	logger.Write("日志测试", "test")
	ReturnSuccess(c, 200, "请求成功", "")
}
