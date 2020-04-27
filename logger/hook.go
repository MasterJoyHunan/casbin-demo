package logger

import (
	"casbindemo/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type MyHook struct{}

func (h *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *MyHook) Fire(entry *logrus.Entry) error {

	fileName := time.Now().Format("2006-01-02")
	fullPath := fmt.Sprintf("%s/%s.log", config.LogConfig.Path, fileName)

	if err := os.MkdirAll(config.LogConfig.Path, os.ModePerm); err != nil {
		log.Panic("创建文件夹错误", err)
	}

	src, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("写入日志文件错误", err)
	}

	//设置输出
	Logger.Out = src
	return nil
}
