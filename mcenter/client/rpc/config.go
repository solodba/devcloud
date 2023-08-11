package rpc

import "context"

// 定义grpc参数结构体
type Config struct {
	address string
	// 添加grpc中间认证参数
	clientID     string
	clientSecret string
}

// Config结构体初始化函数
func NewConfig() *Config {
	return &Config{
		address: "127.0.0.1:8881",
	}
}

// config实现credentials.PerRPCCredentials接口后即可把clientID和clientSecret参数值传递到GRPC客户端中
func (c *Config) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"client_id":     c.clientID,
		"client_secret": c.clientSecret,
	}, nil
}

func (c *Config) RequireTransportSecurity() bool {
	return false
}
