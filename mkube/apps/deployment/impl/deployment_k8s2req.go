package impl

import (
	"github.com/solodba/devcloud/mkube/apps/deployment"
	"github.com/solodba/devcloud/mkube/apps/pod"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *impl) DeploymentK8s2ReqConvert(k8sDeployment *appsv1.Deployment) *deployment.CreateDeploymentRequest {
	return &deployment.CreateDeploymentRequest{
		Base: &deployment.Base{
			Name:      k8sDeployment.Name,
			Namespace: k8sDeployment.Namespace,
			Labels:    i.getDeploymentReqLabels(k8sDeployment.Labels),
			Selector:  i.getDeploymentReqSelector(k8sDeployment.Spec.Selector),
			Replicas:  i.getDeploymentReqReplicas(k8sDeployment.Spec.Replicas),
		},
		Template: i.getDeploymentReqTemplate(k8sDeployment.Spec.Template),
	}
}

func (i *impl) getDeploymentReqLabels(k8sDeploymentLabels map[string]string) []*deployment.ListMapItem {
	deploymentReqLables := make([]*deployment.ListMapItem, 0)
	for k, v := range k8sDeploymentLabels {
		deploymentReqLables = append(deploymentReqLables, &deployment.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return deploymentReqLables
}

func (i *impl) getDeploymentReqSelector(k8sDeploymentSelector *metav1.LabelSelector) []*deployment.ListMapItem {
	deploymentReqSelector := make([]*deployment.ListMapItem, 0)
	if k8sDeploymentSelector == nil {
		return deploymentReqSelector
	}
	if k8sDeploymentSelector.MatchLabels == nil {
		return deploymentReqSelector
	}
	for k, v := range k8sDeploymentSelector.MatchLabels {
		deploymentReqSelector = append(deploymentReqSelector, &deployment.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return deploymentReqSelector
}

func (i *impl) getDeploymentReqReplicas(k8sDeploymentReaplicas *int32) int32 {
	var replicas int32
	if k8sDeploymentReaplicas == nil {
		return replicas
	}
	replicas = *k8sDeploymentReaplicas
	return replicas
}

func (i *impl) getDeploymentReqTemplate(k8sPodTemplateSpec corev1.PodTemplateSpec) *pod.Pod {
	var k8sPod corev1.Pod
	k8sPod.TypeMeta.APIVersion = "v1"
	k8sPod.TypeMeta.Kind = "Pod"
	k8sPod.ObjectMeta = k8sPodTemplateSpec.ObjectMeta
	k8sPod.Spec = k8sPodTemplateSpec.Spec
	podReq := i.podSvc.PodK8s2Req(&k8sPod)
	return podReq
}
