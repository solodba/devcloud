package rpc

// 定义grpc参数结构体
type Config struct {
	address string
}

// Config结构体初始化函数
func NewConfig() *Config {
	return &Config{
		address: "127.0.0.1:8885",
	}
}
