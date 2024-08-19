package db

import (
	"fmt"
	"time"
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

	// 连接到数据库
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
	sqlDB, err := Db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)

	return nil

}
