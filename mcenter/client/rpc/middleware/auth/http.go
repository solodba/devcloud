package auth

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mcenter/apps/endpoint"
	"github.com/solodba/devcloud/mcenter/apps/permission"
	"github.com/solodba/devcloud/mcenter/apps/service"
	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/devcloud/mcenter/client/rpc"
	"github.com/solodba/mcube/logger"
	"github.com/solodba/mcube/response"
)

// httpAuther结构体
type httpAuther struct {
	client *rpc.Client
}

// httpAuther结构体初始化函数
func NewHttpAuther(conf *rpc.McenterGrpcClientConfig) *httpAuther {
	return &httpAuther{
		client: rpc.NewClient(conf),
	}
}

// Http Restful中间件
func (h *httpAuther) AuthFunc(r *restful.Request, w *restful.Response, next *restful.FilterChain) {
	// 构造endpoint结构体
	ep := endpoint.NewDefaultEndpoint()
	route := r.SelectedRoute()
	// 通过grpc获取service id
	serviceReq := service.NewQueryServiceIdByClientIdRequest(h.client.Conf.ClientID)
	serviceIns, err := h.client.NewServiceRPCClient().QueryServiceIdByClientId(r.Request.Context(), serviceReq)
	if err != nil {
		w.WriteEntity(response.NewFail(400, "grpc获取service id错误!"))
		return
	}
	ep.Spec.ServiceId = serviceIns.Meta.Id
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
		tk, err := h.client.NewTokenRPCClient().ValidateToken(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, "令牌校验失败!"))
			return
		}
		r.SetAttribute(token.ATTRIBUTE_TOKEN_KEY, tk)

		// 是否鉴权, 鉴权操作是在用户认证后开始
		if ep.Spec.Perm {
			checkReq := permission.NewCheckPermissionRequest()
			checkReq.UserId = tk.UserId
			checkReq.Namespace = tk.Namespace
			checkReq.ServiceId = ep.Spec.ServiceId
			checkReq.HttpMethod = ep.Spec.Method
			checkReq.HttpPath = ep.Spec.Path
			fmt.Println(checkReq)
			permissionResp, err := h.client.NewPermissionRPCClient().CheckPermission(r.Request.Context(), checkReq)
			if err != nil {
				w.WriteEntity(response.NewFail(500, "权限校验失败!"))
				return
			}
			if !permissionResp.HasPermission {
				w.WriteEntity(response.NewFail(500, "权限校验失败!"))
				return
			}
		}
	}
	next.ProcessFilter(r, w)
}
