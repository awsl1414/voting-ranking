/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-24 17:30:57
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-24 17:34:44
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
	Aid         int    `json:"aId"`         // 活动id
	Ref         string `json:"ref"`         // 编号
	Nickname    string `json:"nickname"`    // 昵称
	Declaration string `json:"declaration"` // 宣言
	Avatar      string `json:"avatar"`      // 头像
	Score       int    `json:"score"`       // 分数
	Phone       string `json:"phone"`       // 电话
	Note        string `json:"note"`        // 备注
}
