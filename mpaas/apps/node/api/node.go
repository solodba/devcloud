package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/node"
	"github.com/solodba/mcube/response"
)

// 查询Node或者Node详情
func (h *handler) QueryOrDescribeNode(r *restful.Request, w *restful.Response) {
	nodeName := r.QueryParameter("name")
	if nodeName == "" {
		req := node.NewQueryNodeRequestFromRestful(r)
		nodeSet, err := h.svc.QueryNode(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, nodeSet))
	} else {
		req := node.NewDescribeNodeRequestFromRestful(r)
		node, err := h.svc.DescribeNode(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, node))
	}
}

// 更新节点标签
func (h *handler) UpdateNodeLabel(r *restful.Request, w *restful.Response) {
	req := node.NewUpdatedLabelRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	_, err = h.svc.UpdateNodeLabel(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, fmt.Sprintf("[%s]节点标签更新成功", req.Name)))
}

// 更新节点污点
func (h *handler) UpdateNodeTaint(r *restful.Request, w *restful.Response) {

}
