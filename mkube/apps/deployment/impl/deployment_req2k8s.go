package impl

import (
	"github.com/solodba/devcloud/mkube/apps/deployment"
	"github.com/solodba/devcloud/mkube/apps/pod"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义Deployment结构体转换成K8S中结构体
func (i *impl) DeploymentReq2K8sConvert(deploymentReq *deployment.CreateDeploymentRequest) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentReq.Base.Name,
			Namespace: deploymentReq.Base.Namespace,
			Labels:    i.getk8sDeploymentLabels(deploymentReq.Base.Labels),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deploymentReq.Base.Replicas,
			Selector: i.getk8sDeploymentSelector(deploymentReq.Base.Selector),
			Template: i.getk8sDeploymentTemplate(deploymentReq.Template),
		},
	}
}

func (i *impl) getk8sDeploymentLabels(deploymentReqLabels []*deployment.ListMapItem) map[string]string {
	k8sDeploymentLables := make(map[string]string)
	for _, label := range deploymentReqLabels {
		k8sDeploymentLables[label.Key] = label.Value
	}
	return k8sDeploymentLables
}

func (i *impl) getk8sDeploymentSelector(deploymentReqSelector []*deployment.ListMapItem) *metav1.LabelSelector {
	var k8sLableSelector metav1.LabelSelector
	k8sLableSelector.MatchLabels = make(map[string]string)
	for _, selector := range deploymentReqSelector {
		k8sLableSelector.MatchLabels[selector.Key] = selector.Value
	}
	return &k8sLableSelector
}

func (i *impl) getk8sDeploymentTemplate(deploymentReqTemplate *pod.Pod) corev1.PodTemplateSpec {
	var podTemplateSpec corev1.PodTemplateSpec
	k8sPod := i.podSvc.PodReq2K8s(deploymentReqTemplate)
	podTemplateSpec.ObjectMeta = k8sPod.ObjectMeta
	podTemplateSpec.Spec = k8sPod.Spec
	return podTemplateSpec
}
