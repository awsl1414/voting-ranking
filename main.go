/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-18 16:59:07
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:34:56
 * @FilePath: /voting-ranking/main.go
 * @Description:
 *
 */
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"voting-ranking/common/config"

	"voting-ranking/pkg/db"
	"voting-ranking/pkg/log"
	"voting-ranking/pkg/redis"
	"voting-ranking/router"

	"github.com/gin-gonic/gin"
)

// @title 活动投票系统API接口文档
// @version 1.0
// @description  活动投票系统API接口文档
// @in header
// @name Authorization
func main() {
	// 加载日志
	log := log.Log()
	// 设置启动模式
	gin.SetMode(config.Config.Server.Model)
	// 路由初始化
	router := router.InitRouter()

	srv := &http.Server{
		Addr:    config.Config.Server.Address, // 服务器监听的地址和端口
		Handler: router,                       // 设置处理请求的路由
	}

	// 启动服务，使用 goroutine 来并行处理
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 如果出现错误且不是服务器关闭的错误，记录错误信息
			log.Infof("listen: %s \n", err)
		}
		log.Infof("listen: %s \n", config.Config.Server.Address)
	}()

	// 创建一个信号通道，用于接收系统信号（如终止信号）
	quit := make(chan os.Signal, 1)
	// 监听操作系统的中断信号（如 Ctrl+C）
	signal.Notify(quit, os.Interrupt)
	// 阻塞，直到接收到中断信号
	<-quit
	// 接收到中断信号后，记录日志并开始服务器关闭流程
	log.Infof("Shutdown Server ...")
	// 创建一个超时上下文，设置为5秒，用于控制服务器的关闭时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 尝试优雅地关闭服务器（完成所有请求后关闭）
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown", err)
	}

	// 记录服务器退出日志
	log.Info("Server exiting")

}

// 初始化连接
func init() {
	// mysql
	db.SetupDBLink()
	// redis
	redis.SetupRedis()
}
