package auth

import (
	"strings"

	"github.com/emicklei/go-restful/v3"
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
func NewHttpAuther(conf *rpc.Config) *httpAuther {
	return &httpAuther{
		client: rpc.NewClient(conf),
	}
}

// Http Restful中间件
func (h *httpAuther) AuthFunc(r *restful.Request, w *restful.Response, next *restful.FilterChain) {
	headTokenValue := r.HeaderParameter(token.AUTH_TOKEN_KEY)
	tklist := strings.Split(headTokenValue, " ")
	if len(tklist) != 2 {
		w.WriteEntity(response.NewFail(400, "令牌格式错误"))
		return
	}
	tk, err := h.client.
		NewTokenRPCClient().
		ValidateToken(r.Request.Context(), token.NewValidateTokenRequest(tklist[1]))
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	logger.L().Info().Msgf("令牌: %s", tk)
	r.SetAttribute(token.ATTRIBUTE_TOKEN_KEY, tk)
	next.ProcessFilter(r, w)
}
