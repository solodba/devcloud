package impl

import (
	"context"
	"strings"

	"github.com/solodba/devcloud/mpaas/apps/rbac"
	"github.com/solodba/devcloud/mpaas/common"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建或者更新ServcieAccount
func (i *impl) CreateOrUpdateServiceAccount(ctx context.Context, in *rbac.ServiceAccount) (*rbac.Response, error) {
	k8sSa := i.RbacSaReq2K8s(in)
	saApi := i.clientSet.CoreV1().ServiceAccounts(k8sSa.Namespace)
	if _, err := saApi.Get(ctx, k8sSa.Name, metav1.GetOptions{}); err != nil {
		// 创建
		if _, err := saApi.Create(ctx, k8sSa, metav1.CreateOptions{}); err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		// 更新
		if _, err := saApi.Update(ctx, k8sSa, metav1.UpdateOptions{}); err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// 删除ServiceAccount
func (i *impl) DeleteServiceAccount(ctx context.Context, in *rbac.DeleteServiceAccountRequest) (*rbac.Response, error) {
	saApi := i.clientSet.CoreV1().ServiceAccounts(in.Namespace)
	if err := saApi.Delete(ctx, in.Name, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}
	return nil, nil
}

// 查询ServiceAccount集合
func (i *impl) GetServiceAccountList(ctx context.Context, in *rbac.GetServiceAccountListRequest) (*rbac.ServiceAccountSet, error) {
	saApi := i.clientSet.CoreV1().ServiceAccounts(in.Namespace)
	var saSet rbac.ServiceAccountSet
	saSet.Items = make([]*rbac.ServiceAccountSetItem, 0)
	if k8sSaList, err := saApi.List(ctx, metav1.ListOptions{}); err != nil {
		return &saSet, err
	} else {
		for _, k8sSa := range k8sSaList.Items {
			saRes := i.RbacK8s2SaRes(k8sSa)
			if strings.Contains(saRes.Name, in.Keyword) {
				saSet.Items = append(saSet.Items, saRes)
			}
		}
		saSet.Total = int64(len(saSet.Items))
		return &saSet, nil
	}
}

// 创建或者更新ClusterRole
func (i *impl) CreateOrUpdateClusterRole(ctx context.Context, in *rbac.Role) (*rbac.Response, error) {
	k8sClusterRole := i.RbacClusterRoleReq2K8s(in)
	roleApi := i.clientSet.RbacV1().ClusterRoles()
	if _, err := roleApi.Get(ctx, k8sClusterRole.Name, metav1.GetOptions{}); err != nil {
		// 创建
		if _, err := roleApi.Create(ctx, k8sClusterRole, metav1.CreateOptions{}); err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		// 更新
		if _, err := roleApi.Update(ctx, k8sClusterRole, metav1.UpdateOptions{}); err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// 删除ClusterRole
func (i *impl) DeleteClusterRole(ctx context.Context, in *rbac.DeleteClusterRoleRequest) (*rbac.Response, error) {
	roleApi := i.clientSet.RbacV1().ClusterRoles()
	if err := roleApi.Delete(ctx, in.Name, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}
	return nil, nil
}

// 查询ClusterRole集合
func (i *impl) GetClusterRoleList(ctx context.Context, in *rbac.GetClusterRoleListRequest) (*rbac.ClusterRoleSet, error) {
	roleApi := i.clientSet.RbacV1().ClusterRoles()
	var crSet rbac.ClusterRoleSet
	crSet.Items = make([]*rbac.ClusterRoleSetItem, 0)
	if k8sClusterRoleList, err := roleApi.List(ctx, metav1.ListOptions{}); err != nil {
		return &crSet, err
	} else {
		for _, k8sClusterRole := range k8sClusterRoleList.Items {
			roleRes := i.RbacK8s2ClusterRoleRes(k8sClusterRole)
			if strings.Contains(roleRes.Name, in.Keyword) {
				crSet.Items = append(crSet.Items, roleRes)
			}
		}
		crSet.Total = int64(len(crSet.Items))
		return &crSet, nil
	}
}

// 查询ClusterRole详情
func (i *impl) GetClusterRoleDetail(ctx context.Context, in *rbac.GetClusterRoleDetailRequest) (*rbac.Role, error) {
	roleApi := i.clientSet.RbacV1().ClusterRoles()
	var clusterRole *rbac.Role
	if k8sClusterRole, err := roleApi.Get(ctx, in.Name, metav1.GetOptions{}); err != nil {
		return nil, err
	} else {
		clusterRole = i.RbacK8s2ClusterRoleResDetail(k8sClusterRole)
		return clusterRole, nil
	}
}

// 创建Role
func (i *impl) CreateOrUpdateRole(ctx context.Context, in *rbac.Role) (*rbac.Response, error) {
	k8sRole := i.RbacRoleReq2K8s(in)
	roleApi := i.clientSet.RbacV1().Roles(k8sRole.Namespace)
	if _, err := roleApi.Get(ctx, k8sRole.Name, metav1.GetOptions{}); err != nil {
		// 创建
		if _, err := roleApi.Create(ctx, k8sRole, metav1.CreateOptions{}); err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		// 更新
		if _, err := roleApi.Update(ctx, k8sRole, metav1.UpdateOptions{}); err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// 删除Role
func (i *impl) DeleteRole(ctx context.Context, in *rbac.DeleteRoleRequest) (*rbac.Response, error) {
	roleApi := i.clientSet.RbacV1().Roles(in.Namespace)
	if err := roleApi.Delete(ctx, in.Name, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}
	return nil, nil
}

// 查询Role集合
func (i *impl) GetRoleList(ctx context.Context, in *rbac.GetRoleListRequest) (*rbac.RoleSet, error) {
	roleApi := i.clientSet.RbacV1().Roles(in.Namespace)
	var roleSet rbac.RoleSet
	roleSet.Items = make([]*rbac.RoleSetItem, 0)
	if k8sRoleList, err := roleApi.List(ctx, metav1.ListOptions{}); err != nil {
		return nil, err
	} else {
		for _, k8sRole := range k8sRoleList.Items {
			roleRes := i.RbacK8s2RoleRes(k8sRole)
			if strings.Contains(roleRes.Name, in.Keyword) {
				roleSet.Items = append(roleSet.Items, roleRes)
			}
		}
		return &roleSet, nil
	}
}

// 查询Role详情
func (i *impl) GetRoleDetail(ctx context.Context, in *rbac.GetRoleDetailRequest) (*rbac.Role, error) {
	roleApi := i.clientSet.RbacV1().Roles(in.Namespace)
	var role *rbac.Role
	if k8sRole, err := roleApi.Get(ctx, in.Name, metav1.GetOptions{}); err != nil {
		return nil, err
	} else {
		role = i.RbacK8s2RoleResDetail(k8sRole)
		return role, nil
	}
}

// 创建RoleBinding
func (i *impl) CreateOrUpdateRb(ctx context.Context, in *rbac.RoleBinding) (*rbac.Response, error) {
	//创建 cluster role binding
	if in.Namespace == "" {
		rbK8sReq := &rbacv1.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      in.Name,
				Namespace: in.Namespace,
				Labels:    common.ToMap(in.Labels),
			},
			Subjects: func(saList []*rbac.ServiceAccount) []rbacv1.Subject {
				subjects := make([]rbacv1.Subject, len(saList))
				for index, item := range saList {
					subjects[index] = rbacv1.Subject{
						Name:      item.Name,
						Kind:      "User",
						Namespace: item.Namespace,
					}
				}
				return subjects
			}(in.Subjects),
			RoleRef: rbacv1.RoleRef{
				Name:     in.RoleRef,
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
			},
		}
		clusterRbApi := i.clientSet.RbacV1().ClusterRoleBindings()
		if cluterRoleSrc, err := clusterRbApi.Get(ctx, in.Name, metav1.GetOptions{}); err != nil {
			_, err := clusterRbApi.Create(ctx, rbK8sReq, metav1.CreateOptions{})
			if err != nil {
				return nil, err
			}
		} else {
			cluterRoleSrc.ObjectMeta.Labels = rbK8sReq.Labels
			cluterRoleSrc.Subjects = rbK8sReq.Subjects
			cluterRoleSrc.RoleRef = rbK8sReq.RoleRef
			_, err := clusterRbApi.Update(ctx, cluterRoleSrc, metav1.UpdateOptions{})
			if err != nil {
				return nil, err
			}
		}
	} else {
		rbK8sReq := &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      in.Name,
				Namespace: in.Namespace,
				Labels:    common.ToMap(in.Labels),
			},
			Subjects: func(saList []*rbac.ServiceAccount) []rbacv1.Subject {
				subjects := make([]rbacv1.Subject, len(saList))
				for index, item := range saList {
					subjects[index] = rbacv1.Subject{
						Name:      item.Name,
						Kind:      "User",
						Namespace: item.Namespace,
					}
				}
				return subjects
			}(in.Subjects),
			RoleRef: rbacv1.RoleRef{
				Name:     in.RoleRef,
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "Role",
			},
		}
		rbApi := i.clientSet.RbacV1().RoleBindings(rbK8sReq.Namespace)
		if rbK8sSrc, err := rbApi.Get(ctx, in.Name, metav1.GetOptions{}); err != nil {
			_, err = rbApi.Create(ctx, rbK8sReq, metav1.CreateOptions{})
			if err != nil {
				return nil, err
			}
		} else {
			rbK8sSrc.ObjectMeta.Labels = rbK8sReq.Labels
			rbK8sSrc.Subjects = rbK8sReq.Subjects
			rbK8sSrc.RoleRef = rbK8sReq.RoleRef
			_, err = rbApi.Update(ctx, rbK8sSrc, metav1.UpdateOptions{})
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

// 删除RoleBinding
func (i *impl) DeleteRb(ctx context.Context, in *rbac.DeleteRbRequest) (*rbac.Response, error) {
	if in.Namespace != "" {
		err := i.clientSet.RbacV1().RoleBindings(in.Namespace).
			Delete(ctx, in.Name, metav1.DeleteOptions{})
		if err != nil {
			return nil, err
		}
	} else {
		err := i.clientSet.RbacV1().ClusterRoleBindings().
			Delete(ctx, in.Name, metav1.DeleteOptions{})
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// 查看RoleBinding列表
func (i *impl) GetRbList(ctx context.Context, in *rbac.GetRbListRequest) (*rbac.RoleBindingSet, error) {
	var rbSet rbac.RoleBindingSet
	rbSet.Items = make([]*rbac.RoleBindingSetItem, 0)
	if in.Namespace != "" {
		list, err := i.clientSet.RbacV1().RoleBindings(in.Namespace).
			List(ctx, metav1.ListOptions{})
		if err != nil {
			return &rbSet, err
		}
		for _, item := range list.Items {
			if !strings.Contains(item.Name, in.Keyword) {
				continue
			}
			rbSet.Items = append(rbSet.Items, &rbac.RoleBindingSetItem{
				Name:      item.Name,
				Namespace: item.Namespace,
				Age:       item.CreationTimestamp.Unix(),
			})
		}
	} else {
		list, err := i.clientSet.RbacV1().ClusterRoleBindings().
			List(ctx, metav1.ListOptions{})
		if err != nil {
			return &rbSet, err
		}
		for _, item := range list.Items {
			if !strings.Contains(item.Name, in.Keyword) {
				continue
			}
			rbSet.Items = append(rbSet.Items, &rbac.RoleBindingSetItem{
				Name:      item.Name,
				Namespace: item.Namespace,
				Age:       item.CreationTimestamp.Unix(),
			})
		}
	}
	return &rbSet, nil
}

// 查询RoleBinding详情
func (i *impl) GetRbDetail(ctx context.Context, in *rbac.GetRbDetailRequest) (*rbac.RoleBinding, error) {
	rbReq := rbac.RoleBinding{}
	if in.Namespace != "" {
		rbK8s, err := i.clientSet.RbacV1().RoleBindings(in.Namespace).
			Get(ctx, in.Name, metav1.GetOptions{})
		if err != nil {
			return &rbReq, err
		}
		rbReq.Name = rbK8s.Name
		rbReq.Namespace = rbK8s.Namespace
		rbReq.Labels = common.ToList(rbK8s.Labels)
		rbReq.RoleRef = rbK8s.RoleRef.Name
		rbReq.Subjects = func(subjects []rbacv1.Subject) []*rbac.ServiceAccount {
			saList := make([]*rbac.ServiceAccount, len(subjects))
			for i, subject := range subjects {
				saList[i] = &rbac.ServiceAccount{
					Name:      subject.Name,
					Namespace: subject.Namespace,
				}
			}
			return saList
		}(rbK8s.Subjects)
	} else {
		rbK8s, err := i.clientSet.RbacV1().ClusterRoleBindings().
			Get(ctx, in.Name, metav1.GetOptions{})
		if err != nil {
			return &rbReq, err
		}
		rbReq.Name = rbK8s.Name
		rbReq.Namespace = rbK8s.Namespace
		rbReq.Labels = common.ToList(rbK8s.Labels)
		rbReq.RoleRef = rbK8s.RoleRef.Name
		rbReq.Subjects = func(subjects []rbacv1.Subject) []*rbac.ServiceAccount {
			saList := make([]*rbac.ServiceAccount, len(subjects))
			for i, subject := range subjects {
				saList[i] = &rbac.ServiceAccount{
					Name:      subject.Name,
					Namespace: subject.Namespace,
				}
			}
			return saList
		}(rbK8s.Subjects)
	}
	return &rbReq, nil
}
