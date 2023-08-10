package rest

// 配置结构体
type Config struct {
	url      string
	username string
	password string
}

// Config初始化函数
func NewConfig() *Config {
	return &Config{
		url:      "http://localhost:8884/mcenter/api/v1/",
		username: "admin",
		password: "admin",
	}
}
