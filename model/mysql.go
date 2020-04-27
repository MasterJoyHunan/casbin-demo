package model

import (
	"casbindemo/config"
	"casbindemo/logger"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	Db  *gorm.DB
	err error
)

func Setup() {
	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlConfig.User,
		config.MysqlConfig.Pwd,
		config.MysqlConfig.Host,
		config.MysqlConfig.Port,
		config.MysqlConfig.Db))
	if err != nil {
		log.Panic(fmt.Sprintf("连接数据库错误 ：%s", err))
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.MysqlConfig.Prefix + defaultTableName
	}
	Db.SingularTable(true)
	Db.LogMode(true)
	Db.SetLogger(logger.Logger)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(10)
	AutoMigrate()
}

// 自动创建表结构
func AutoMigrate() {
	Db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT 'casbin 策略'").
		AutoMigrate(&CasbinRule{})
}
