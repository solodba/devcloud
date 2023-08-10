package rest

import "github.com/infraboard/mcube/client/rest"

// restful客户端结构体
type Client struct {
	c *rest.RESTClient
}

// Client结构体初始化函数
func NewClient(conf *Config) *Client {
	r := rest.NewRESTClient()
	r.SetBaseURL(conf.url)
	r.SetBasicAuth(conf.username, conf.password)
	return &Client{
		c: r,
	}
}
