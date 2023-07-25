package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/statefulset"
	appsv1 "k8s.io/api/apps/v1"
)

// k8s中结构体转换成自定义结构体
func (i *impl) StatefulSetlK8s2ResConvert(k8sStatefulSet appsv1.StatefulSet) *statefulset.StatefulSetSetItem {
	return &statefulset.StatefulSetSetItem{
		Name:      k8sStatefulSet.Name,
		Namespace: k8sStatefulSet.Namespace,
		Ready:     k8sStatefulSet.Status.ReadyReplicas,
		Replicas:  k8sStatefulSet.Status.Replicas,
		Age:       k8sStatefulSet.CreationTimestamp.Unix(),
	}
}
