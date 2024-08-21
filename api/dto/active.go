/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-21 23:02:53
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-21 23:21:10
 * @FilePath: /voting-ranking/api/dto/active.go
 * @Description:
 *
 */
package dto

type ActivityDto struct {
	Name string `json:"name" validata:"required"` // 活动名
}
