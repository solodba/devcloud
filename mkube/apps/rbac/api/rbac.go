package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/rbac"
	"github.com/solodba/mcube/response"
)

// 创建ServiceAccount
func (h *handler) CreateOrUpdateServiceAccount(r *restful.Request, w *restful.Response) {
	var saReq rbac.ServiceAccount
	if err := r.ReadEntity(&saReq); err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	if _, err := h.svc.CreateOrUpdateServiceAccount(r.Request.Context(), &saReq); err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, "创建ServiceAccount成功!"))
}

// 删除ServiceAccount
func (h *handler) DeleteServiceAccount(r *restful.Request, w *restful.Response) {
	var req rbac.DeleteServiceAccountRequest
	req.Namespace = r.PathParameter("namespace")
	req.Name = r.PathParameter("name")
	if _, err := h.svc.DeleteServiceAccount(r.Request.Context(), &req); err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, "删除ServiceAccount成功!"))
}

// 获取ServiceAccount列表
func (h *handler) GetServiceAccountList(r *restful.Request, w *restful.Response) {
	var req rbac.GetServiceAccountListRequest
	req.Namespace = r.PathParameter("namespace")
	req.Keyword = r.QueryParameter("keyword")
	saList, err := h.svc.GetServiceAccountList(r.Request.Context(), &req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, saList))
}

// 创建Role
func (h *handler) CreateOrUpdateRole(r *restful.Request, w *restful.Response) {
	var roleReq rbac.Role
	if err := r.ReadEntity(&roleReq); err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	if roleReq.Namespace == "" {
		// 创建cluster role
		if _, err := h.svc.CreateOrUpdateClusterRole(r.Request.Context(), &roleReq); err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		} else {
			w.WriteEntity(response.NewSuccess(200, "创建Cluster Role成功!"))
		}
	} else {
		// 创建namespace role
		if _, err := h.svc.CreateOrUpdateRole(r.Request.Context(), &roleReq); err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		} else {
			w.WriteEntity(response.NewSuccess(200, "创建Namespace Role成功!"))
		}
	}
}

// 删除Role
func (h *handler) DeleteRole(r *restful.Request, w *restful.Response) {
	namespace := r.PathParameter("namespace")
	name := r.PathParameter("name")
	if namespace == "" {
		// 删除Cluster role
		var deleteClusterRoleReq rbac.DeleteClusterRoleRequest
		deleteClusterRoleReq.Name = name
		if _, err := h.svc.DeleteClusterRole(r.Request.Context(), &deleteClusterRoleReq); err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		} else {
			w.WriteEntity(response.NewSuccess(200, "删除Cluster Role成功!"))
		}
	} else {
		// 删除namespace role
		var deleteRoleReq rbac.DeleteRoleRequest
		deleteRoleReq.Name = name
		deleteRoleReq.Namespace = namespace
		if _, err := h.svc.DeleteRole(r.Request.Context(), &deleteRoleReq); err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		} else {
			w.WriteEntity(response.NewSuccess(200, "删除Namespace Role成功!"))
		}
	}
}

// 获取Role详情或者列表
func (h *handler) GetRoleDetailOrList(r *restful.Request, w *restful.Response) {
	namespace := r.QueryParameter("namespace")
	name := r.QueryParameter("name")
	keyword := r.QueryParameter("keyword")
	if namespace != "" {
		if name != "" {
			var getRoleDetailRequest rbac.GetRoleDetailRequest
			getRoleDetailRequest.Namespace = namespace
			getRoleDetailRequest.Name = name
			if role, err := h.svc.GetRoleDetail(r.Request.Context(), &getRoleDetailRequest); err != nil {
				w.WriteEntity(response.NewFail(500, err.Error()))
				return
			} else {
				w.WriteEntity(response.NewSuccess(200, role))
			}
			return
		}
		var getRoleListRequest rbac.GetRoleListRequest
		getRoleListRequest.Namespace = namespace
		getRoleListRequest.Keyword = keyword
		if roleList, err := h.svc.GetRoleList(r.Request.Context(), &getRoleListRequest); err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		} else {
			w.WriteEntity(response.NewSuccess(200, roleList))
		}
	} else {
		if name != "" {
			var getClusterRoleDetailRequest rbac.GetClusterRoleDetailRequest
			getClusterRoleDetailRequest.Name = name
			if clusterRole, err := h.svc.GetClusterRoleDetail(r.Request.Context(), &getClusterRoleDetailRequest); err != nil {
				w.WriteEntity(response.NewFail(500, err.Error()))
				return
			} else {
				w.WriteEntity(response.NewSuccess(200, clusterRole))
			}
			return
		}
		var getClusterRoleListRequest rbac.GetClusterRoleListRequest
		getClusterRoleListRequest.Keyword = keyword
		if clusterRoleList, err := h.svc.GetClusterRoleList(r.Request.Context(), &getClusterRoleListRequest); err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		} else {
			w.WriteEntity(response.NewSuccess(200, clusterRoleList))
		}
	}
}

func (h *handler) CreateOrUpdateRb(r *restful.Request, w *restful.Response) {
	var rbReq rbac.RoleBinding
	if err := r.ReadEntity(&rbReq); err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	_, err := h.svc.CreateOrUpdateRb(r.Request.Context(), &rbReq)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, "创建RoleBinding成功!"))
}

func (h *handler) GetRbDetailOrList(r *restful.Request, w *restful.Response) {
	namespace := r.QueryParameter("namespace")
	name := r.QueryParameter("name")
	keyword := r.QueryParameter("keyword")
	if name != "" {
		var getRbDetailRequest rbac.GetRbDetailRequest
		getRbDetailRequest.Namespace = namespace
		getRbDetailRequest.Name = name
		detail, err := h.svc.GetRbDetail(r.Request.Context(), &getRbDetailRequest)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, detail))
	} else {
		var getRbListRequest rbac.GetRbListRequest
		getRbListRequest.Namespace = namespace
		getRbListRequest.Keyword = keyword
		list, err := h.svc.GetRbList(r.Request.Context(), &getRbListRequest)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, list))
	}
}

func (h *handler) DeleteRb(r *restful.Request, w *restful.Response) {
	var req rbac.DeleteRbRequest
	req.Namespace = r.QueryParameter("namespace")
	req.Name = r.QueryParameter("name")
	_, err := h.svc.DeleteRb(r.Request.Context(), &req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, "RoleBinding删除成功!"))
}
