package impl

import (
	"github.com/solodba/devcloud/mkube/apps/daemonset"
	appsv1 "k8s.io/api/apps/v1"
)

func (i *impl) DaemonSetK8s2ResConvert(k8sDaemonSet appsv1.DaemonSet) *daemonset.DaemonSetListItem {
	return &daemonset.DaemonSetListItem{
		Name:      k8sDaemonSet.Name,
		Namespace: k8sDaemonSet.Namespace,
		Desired:   k8sDaemonSet.Status.DesiredNumberScheduled,
		Current:   k8sDaemonSet.Status.CurrentNumberScheduled,
		Ready:     k8sDaemonSet.Status.NumberReady,
		UpToDate:  k8sDaemonSet.Status.UpdatedNumberScheduled,
		Available: k8sDaemonSet.Status.NumberAvailable,
		Age:       k8sDaemonSet.CreationTimestamp.Unix(),
	}
}
