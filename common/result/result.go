/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 10:31:50
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-19 15:13:12
 * @FilePath: /voting-ranking/common/result/result.go
 * @Description:
 *
 */
package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result 统一返回结果
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 返回成功的 JSON 响应
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{
		Code:    int(ApiCode.SUCCESS),
		Message: ApiCode.GetMessage(ApiCode.SUCCESS),
		Data:    data,
	}

	c.JSON(http.StatusOK, res)
}

// Failed 返回失败的 JSON 响应
func Failed(c *gin.Context, code int, message string) {
	res := Result{
		Code:    code,
		Message: message,
		Data:    gin.H{},
	}
	c.JSON(http.StatusOK, res)
}
