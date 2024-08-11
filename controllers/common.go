/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 21:45:40
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-11 22:16:05
 * @FilePath: /hnuahe-presentation-voting-ranking/controllers/common.go
 * @Description:
 *
 */

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data}
	c.JSON(http.StatusOK, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}, data interface{}) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}
