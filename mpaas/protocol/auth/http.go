package auth

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mcenter/apps/endpoint"
	"github.com/solodba/mcube/logger"
)

// httpAuther认证结构体
type httpAuther struct {
}

// httpAuther初始化函数
func NewHttpAuther() *httpAuther {
	return &httpAuther{}
}

// type FilterFunction func(*Request, *Response, *FilterChain)
// 实现上述函数就是http中间件
func (a *httpAuther) AuthFunc(r *restful.Request, w *restful.Response, next *restful.FilterChain) {
	// 构造endpoint结构体
	ep := endpoint.NewDefaultEndpoint()
	route := r.SelectedRoute()
	// 服务的service id可以gprc去获取, 这里写死了
	ep.Spec.ServiceId = "cjauuq4fd1fkek1bmfpg"
	ep.Spec.Method = route.Method()
	ep.Spec.Operation = route.Operation()
	ep.Spec.Path = route.Path()
	isAuth := route.Metadata()["auth"]
	if isAuth != nil {
		ep.Spec.Auth = isAuth.(bool)
	}
	logger.L().Info().Msgf("%v", ep.Spec)
	next.ProcessFilter(r, w)
}
