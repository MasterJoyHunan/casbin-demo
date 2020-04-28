package config

// app 配置
type Application struct {
	Env  string
	Host string
	Port int
}

// mysql 数据库配置
type Mysql struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	Db     string
	Prefix string
}

// 日志配置
type Log struct {
	Path  string
	Level string
}

// casbin配置
type Casbin struct {
	Path string
}
