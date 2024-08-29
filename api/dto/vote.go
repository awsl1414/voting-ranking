/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-29 10:20:41
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-29 10:21:37
 * @FilePath: /voting-ranking/api/dto/vote.go
 * @Description:
 *
 */
package dto

// VoteDto 投票参数
type VoteDto struct {
	PlayerId int `json:"playerId" validate:"required"` // 用户名
	UserId   int `json:"userId" validate:"required"`   // 密码
}
