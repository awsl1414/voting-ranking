/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 17:30:57
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-25 23:32:47
 * @FilePath: /voting-ranking/api/dto/player.go
 * @Description:
 *
 */
package dto

// 获取选手列表
type PlayerListDto struct {
	Aid int `json:"aid" validate:"required"` // 活动id
}

// 新增选手
type AddPlayerDto struct {
	Aid         int    `json:"aid" validate:"required"`         // 活动id
	Ref         string `json:"ref" validate:"required"`         // 编号
	Nickname    string `json:"nickname" validate:"required"`    // 昵称
	Declaration string `json:"declaration" validate:"required"` // 宣言
	Avatar      string `json:"avatar" validate:"required"`      // 头像
	Score       int    `json:"score" validate:"required"`       // 分数
	Phone       string `json:"phone" validate:"required"`       // 电话
	Note        string `json:"note" validate:"required"`        // 备注
}
