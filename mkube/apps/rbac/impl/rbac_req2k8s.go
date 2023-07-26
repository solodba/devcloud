package impl

import (
	"github.com/solodba/devcloud/mkube/apps/rbac"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *impl) RbacSaReq2K8s(saReq *rbac.ServiceAccount) *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      saReq.Name,
			Namespace: saReq.Namespace,
			Labels:    i.getK8sSaLabels(saReq.Labels),
		},
	}
}

func (i *impl) getK8sSaLabels(saReqLabels []*rbac.ListMapItem) map[string]string {
	k8sSaLables := make(map[string]string)
	for _, label := range saReqLabels {
		k8sSaLables[label.Key] = label.Value
	}
	return k8sSaLables
}

func (i *impl) RbacClusterRoleReq2K8s(roleReq *rbac.Role) *rbacv1.ClusterRole {
	return &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleReq.Name,
			Namespace: roleReq.Namespace,
			Labels:    i.getK8sRoleLabels(roleReq.Labels),
		},
		Rules: i.getK8sRoleRules(roleReq.Rules),
	}
}

// Rules转换
func (i *impl) getK8sRoleRules(roleReqRules []*rbac.PolicyRule) []rbacv1.PolicyRule {
	k8sRbacPolicyRules := make([]rbacv1.PolicyRule, 0)
	for _, item := range roleReqRules {
		k8sRbacPolicyRules = append(k8sRbacPolicyRules, rbacv1.PolicyRule{
			Verbs:         item.Verbs,
			APIGroups:     item.APIGroups,
			Resources:     item.Resources,
			ResourceNames: item.ResourceNames,
		})
	}
	return k8sRbacPolicyRules
}

func (i *impl) getK8sRoleLabels(roleReqLables []*rbac.ListMapItem) map[string]string {
	k8sRoleLables := make(map[string]string)
	for _, label := range roleReqLables {
		k8sRoleLables[label.Key] = label.Value
	}
	return k8sRoleLables
}

func (i *impl) RbacRoleReq2K8s(roleReq *rbac.Role) *rbacv1.Role {
	return &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleReq.Name,
			Namespace: roleReq.Namespace,
			Labels:    i.getK8sRoleLabels(roleReq.Labels),
		},
		Rules: i.getK8sRoleRules(roleReq.Rules),
	}
}
