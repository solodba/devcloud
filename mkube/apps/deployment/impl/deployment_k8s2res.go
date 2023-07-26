package impl

import (
	"github.com/solodba/devcloud/mkube/apps/deployment"
	appsv1 "k8s.io/api/apps/v1"
)

func (i *impl) DeploymentK8s2ResConvert(k8sDeployment *appsv1.Deployment) *deployment.DeploymentSetItem {
	return &deployment.DeploymentSetItem{
		Name:      k8sDeployment.Name,
		Namespace: k8sDeployment.Namespace,
		Replicas:  i.getDeploymentResReplicas(k8sDeployment.Spec.Replicas),
		Ready:     k8sDeployment.Status.ReadyReplicas,
		UpToDate:  k8sDeployment.Status.UpdatedReplicas,
		Available: k8sDeployment.Status.AvailableReplicas,
		Age:       k8sDeployment.CreationTimestamp.Unix(),
	}
}

func (i *impl) getDeploymentResReplicas(k8sDeploymentReplicas *int32) int32 {
	var replicas int32
	if k8sDeploymentReplicas == nil {
		return replicas
	}
	replicas = *k8sDeploymentReplicas
	return replicas
}
