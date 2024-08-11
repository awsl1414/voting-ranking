/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 23:21:19
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-12 00:06:35
 * @FilePath: /hnuahe-presentation-voting-ranking/dao/dao.go
 * @Description:
 *
 */
package dao

import (
	"time"
	"voting-ranking/config"
	"voting-ranking/pkg/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

// init 函数会在包初始化时自动执行，用于初始化数据库连接
func init() {
	// 使用 GORM 打开 MySQL 数据库连接
	Db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	// 如果连接失败，记录错误日志
	if err != nil {
		logger.Error(logrus.Fields{"mariadb connect error": err.Error()})
	}
	// 检查数据库连接是否有错误
	if Db.Error != nil {
		logger.Error(logrus.Fields{"database error": Db.Error})
	}
	// 获取底层的 SQL DB 连接对象
	sqlDB, _ := Db.DB()

	// 设置连接池的最大空闲连接数
	sqlDB.SetMaxIdleConns(10)

	// 设置连接池的最大连接数
	sqlDB.SetMaxOpenConns(100)

	// 设置连接的最大存活时间
	sqlDB.SetConnMaxLifetime(time.Hour)

}
