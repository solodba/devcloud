package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/mcube/response"
)

// 查询集群信息
func (h *handler) QueryCluster(r *restful.Request, w *restful.Response) {
	AttrTk := r.Attribute(token.ATTRIBUTE_TOKEN_KEY)
	tk := token.NewToken()
	if AttrTk != nil {
		tk = AttrTk.(*token.Token)
	}
	w.WriteEntity(response.NewSuccess(200, tk))
}
