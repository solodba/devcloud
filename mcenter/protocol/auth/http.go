package auth

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/mcube/apps"
)

// httpAuther认证结构体
type httpAuther struct {
	svc token.Service
}

// httpAuther初始化函数
func NewHttpAuther() *httpAuther {
	return &httpAuther{
		svc: apps.GetInternalApp(token.AppName).(token.Service),
	}
}

// type FilterFunction func(*Request, *Response, *FilterChain)
// 实现上述函数就是http中间件
func (a *httpAuther) AuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 获取http认证token信息
	// tokenHeader := req.HeaderParameter(token.AUTH_TOKEN_KEY)
	// tokenHeaderList := strings.Split(tokenHeader, " ")
}
