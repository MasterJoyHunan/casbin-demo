package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	MysqlConfig       Mysql
	LogConfig         Log
	ApplicationConfig Application
	CasbinConfig      Casbin
)

func Setup() {
	viper.SetConfigFile(Path)
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("读取配置文件错误", err)
	}

	if err := viper.UnmarshalKey("mysql", &MysqlConfig); err != nil {
		log.Panic("数据库配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("log", &LogConfig); err != nil {
		log.Panic("日志配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("application", &ApplicationConfig); err != nil {
		log.Panic("APP配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("casbin", &CasbinConfig); err != nil {
		log.Panic("casbin配置文件格式错误", err)
	}
}
