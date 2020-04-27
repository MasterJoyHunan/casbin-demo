package config

import (
	"flag"
	"os"
)

var Path string

func init() {
	dir, _ := os.Getwd()
	flag.StringVar(&Path, "config", dir + "/config.yaml", "配置文件路径")
}
