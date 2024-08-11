package controllers

import "github.com/gin-gonic/gin"

type Test struct{}

func (t Test) Hello(c *gin.Context) {
	ReturnSuccess(c, 200, "请求成功", "hello")
}
