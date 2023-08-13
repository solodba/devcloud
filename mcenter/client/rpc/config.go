package rpc

import "context"

// 定义grpc参数结构体
type McenterGrpcClientConfig struct {
	Address string `json:"address" toml:"address" env:"MCENTER_GRPC_CLIENT_ADDRESS"`
	// 添加grpc中间认证参数
	ClientID     string `json:"clientid" toml:"clientid" env:"MCENTER_GRPC_CLIENT_ID"`
	ClientSecret string `json:"clientsecret" toml:"clientsecret" env:"MCENTER_GRPC_CLIENT_SECRET"`
}

// Config结构体初始化函数
func NewMcenterGrpcClientConfig() *McenterGrpcClientConfig {
	return &McenterGrpcClientConfig{
		Address: "127.0.0.1:8881",
	}
}

// config实现credentials.PerRPCCredentials接口后即可把clientID和clientSecret参数值传递到GRPC客户端中
func (c *McenterGrpcClientConfig) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
	}, nil
}

func (c *McenterGrpcClientConfig) RequireTransportSecurity() bool {
	return false
}
