package logger

import (
	"casbindemo/config"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var Logger = logrus.New()

func Setup() {

	//写入文件
	if config.LogConfig.Path == "" {
		config.LogConfig.Path, _ = os.Getwd()
		config.LogConfig.Path += "/log"
	}

	level, err := logrus.ParseLevel(config.LogConfig.Level)
	if err != nil {
		log.Panic("日志level格式设置错误", err)
	}
	Logger.SetLevel(level)

	//设置日志格式
	Logger.SetFormatter(&logrus.JSONFormatter{})
	// Logger.SetFormatter(&logrus.TextFormatter{}) // 默认格式 无需显示设置

	// 取消线程安全
	Logger.SetNoLock()

	// 自定义HOOK
	Logger.AddHook(&MyHook{})
}
