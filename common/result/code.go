/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 10:12:51
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-19 14:57:47
 * @FilePath: /voting-ranking/common/result/code.go
 * @Description:
 *
 */

// 状态码

package result

// COdes 定义状态
type Codes struct {
	SUCCESS  uint
	FAILED   uint
	Message  map[uint]string
	REQUIRED uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:  200,
	FAILED:   501,
	REQUIRED: 400,
}

// 状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:  "成功",
		ApiCode.FAILED:   "失败",
		ApiCode.REQUIRED: "缺少必要参数",
	}
}

// 外部调用
func (c Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
