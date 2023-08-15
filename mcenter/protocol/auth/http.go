package auth

import (
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mcenter/apps/endpoint"
	"github.com/solodba/devcloud/mcenter/apps/permission"
	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	"github.com/solodba/mcube/response"
)

// httpAuther认证结构体
type httpAuther struct {
	tokenSvc token.Service
}

// httpAuther初始化函数
func NewHttpAuther() *httpAuther {
	return &httpAuther{
		tokenSvc: apps.GetInternalApp(token.AppName).(token.Service),
	}
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
	isPerm := route.Metadata()["perm"]
	if isPerm != nil {
		ep.Spec.Perm = isPerm.(bool)
	}
	logger.L().Info().Msgf("%v", ep.Spec)
	// 开启认证才认证
	if ep.Spec.Auth {
		// 获取http认证token信息
		tokenHeader := r.HeaderParameter(token.AUTH_TOKEN_KEY)
		tokenHeaderList := strings.Split(tokenHeader, " ")
		if len(tokenHeaderList) != 2 {
			w.WriteEntity(response.NewFail(400, "令牌格式错误!"))
			return
		}
		req := token.NewValidateTokenRequest(tokenHeaderList[1])
		// mcenter的http中间件通过内部服务校验令牌
		tk, err := a.tokenSvc.ValidateToken(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, "令牌校验失败!"))
			return
		}
		r.SetAttribute(token.ATTRIBUTE_TOKEN_KEY, tk)

		// 是否鉴权, 鉴权操作是在用户认证后开始
		if ep.Spec.Perm {
			checkReq := permission.NewCheckPermissionRequest()
			checkReq.ServiceId = ep.Spec.ServiceId

		}
	}
	next.ProcessFilter(r, w)
}
