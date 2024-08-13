/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 21:59:40
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-13 12:17:48
 * @FilePath: /voting-ranking/main.go
 * @Description:
 *
 */
package main

import (
	"voting-ranking/dao"
	"voting-ranking/models"
	"voting-ranking/router"
)

func main() {

	// 数据库迁移
	if err := dao.Db.AutoMigrate(&models.User{}, &models.Player{}); err != nil {
		panic("数据库迁移失败")
	}

	router := router.Router()

	router.Run("127.0.0.1:8888")
}
