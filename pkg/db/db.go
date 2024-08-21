/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 16:54:03
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-20 17:50:45
 * @FilePath: /voting-ranking/pkg/db/db.go
 * @Description:
 *
 */
package db

import (
	"fmt"
	"time"
	"voting-ranking/api/model"
	"voting-ranking/common/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

// 数据库初始化
func SetupDBLink() error {
	var err error
	var dbConfig = config.Config.Db

	// 初始化不带数据库名的 DSN
	dsnNoDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Charset,
	)

	Db, err = gorm.Open(mysql.Open(dsnNoDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 设置 Gorm 的日志记录级别为“Silent”。
	})
	if err != nil {
		panic(err)
	}

	// 检查数据库是否存在
	// err = Db.Exec("USE `" + dbConfig.Db + "`").Error
	// if err != nil {}
	var dbExists bool
	Db.Raw("SELECT EXISTS(SELECT 1 FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?)", dbConfig.Db).Scan(&dbExists)
	if !dbExists {
		// 数据库不存在，创建数据库
		createDBSQL := fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", dbConfig.Db)
		if err := Db.Exec(createDBSQL).Error; err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
		fmt.Println("Database created:", dbConfig.Db)

		// 确保（远程）数据库创建完毕
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Database already exists:", dbConfig.Db)
	}

	// 重新连接到指定数据库
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)
	Db, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	// 配置数据库连接池
	sqlDB, err := Db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)

	// 检查并创建表
	tables := map[string]interface{}{
		"user": &model.User{},
	}
	for tableName, tableModel := range tables {
		var tableExists bool
		query := "SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema = ? AND table_name = ?)"
		err = Db.Raw(query, dbConfig.Db, tableName).Scan(&tableExists).Error
		if err != nil {
			return fmt.Errorf("failed to check table '%s' existence: %w", tableName, err)
		}

		if !tableExists {
			fmt.Printf("Table '%s' does not exist, creating...", tableName)
			err := Db.AutoMigrate(tableModel)
			if err != nil {
				return fmt.Errorf("failed to create table '%s': %w", tableName, err)
			}
		} else {
			fmt.Printf("Table '%s' already exists, skipping creation.", tableName)
		}
	}
	return nil

}
