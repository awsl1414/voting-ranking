/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 21:59:40
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-11 22:21:06
 * @FilePath: /hnuahe-presentation-voting-ranking/main.go
 * @Description:
 *
 */
package main

import "voting-ranking/router"

func main() {
	router := router.Router()

	router.Run("127.0.0.1:8888")
}
