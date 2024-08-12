/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 21:45:40
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-12 16:40:45
 * @FilePath: /hnuahe-presentation-voting-ranking/controllers/common.go
 * @Description:
 *
 */

package controllers

import (
	"crypto/md5"
	"encoding/hex"
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

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}

// md5加密
func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
