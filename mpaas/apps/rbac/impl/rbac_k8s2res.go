package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/rbac"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// 查询转换
func (i *impl) RbacK8s2SaRes(k8sSa corev1.ServiceAccount) *rbac.ServiceAccountSetItem {
	return &rbac.ServiceAccountSetItem{
		Name:      k8sSa.Name,
		Namespace: k8sSa.Namespace,
		Age:       k8sSa.CreationTimestamp.Unix(),
	}
}

func (i *impl) RbacK8s2ClusterRoleRes(k8sRole rbacv1.ClusterRole) *rbac.ClusterRoleSetItem {
	return &rbac.ClusterRoleSetItem{
		Name:      k8sRole.Name,
		Namespace: k8sRole.Namespace,
		Age:       k8sRole.CreationTimestamp.Unix(),
	}
}

func (i *impl) RbacK8s2RoleRes(k8sRole rbacv1.Role) *rbac.RoleSetItem {
	return &rbac.RoleSetItem{
		Name:      k8sRole.Name,
		Namespace: k8sRole.Namespace,
		Age:       k8sRole.CreationTimestamp.Unix(),
	}
}

func (i *impl) RbacK8s2ClusterRoleResDetail(clusterRole *rbacv1.ClusterRole) *rbac.Role {
	return &rbac.Role{
		Name:      clusterRole.Name,
		Namespace: clusterRole.Namespace,
		Labels:    i.getClusterRoleLables(clusterRole.Labels),
		Rules:     i.getRoleRules(clusterRole.Rules),
	}
}

// Rules转换
func (i *impl) getRoleRules(k8sRoleRules []rbacv1.PolicyRule) []*rbac.PolicyRule {
	rolePolicyRules := make([]*rbac.PolicyRule, 0)
	for _, item := range k8sRoleRules {
		rolePolicyRules = append(rolePolicyRules, &rbac.PolicyRule{
			Verbs:         item.Verbs,
			APIGroups:     item.APIGroups,
			Resources:     item.Resources,
			ResourceNames: item.ResourceNames,
		})
	}
	return rolePolicyRules
}

func (i *impl) getClusterRoleLables(k8sClusterRoleLables map[string]string) []*rbac.ListMapItem {
	clusterRoleLables := make([]*rbac.ListMapItem, 0)
	for k, v := range k8sClusterRoleLables {
		clusterRoleLables = append(clusterRoleLables, &rbac.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return clusterRoleLables
}

func (i *impl) RbacK8s2RoleResDetail(clusterRole *rbacv1.Role) *rbac.Role {
	return &rbac.Role{
		Name:      clusterRole.Name,
		Namespace: clusterRole.Namespace,
		Labels:    i.getClusterRoleLables(clusterRole.Labels),
		Rules:     i.getRoleRules(clusterRole.Rules),
	}
}
