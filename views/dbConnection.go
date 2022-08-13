package views

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
)

var db *gorm.DB

func init() {
	file, err := ioutil.ReadFile("./DBConfig.txt")
	if err != nil {
		log.Fatalf("读取数据库配置文件失败: %v", err)
	}
	dsn := string(file)
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&")
	fmt.Println(dsn)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为Info, 即打印所有SQL语句
		// Logger:logger.Default.LogMode(logger.Warn), // 只打印慢查询, 默认的SlowThreshold为200ms
	})
	if err != nil {
		// 关闭数据库连接，打印错误信息
		log.Fatalf("数据库连接错误: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("数据库初始化错误: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)  // 设置连接池最大空闲连接数
	sqlDB.SetMaxOpenConns(200) // 设置最大打开连接数
	// sqlDB.SetConnMaxLifetime() // 设置最大连接时长
}
