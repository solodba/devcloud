package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/job"
	"github.com/solodba/devcloud/mpaas/apps/pod"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义job结构体转换成k8s中结构体
func (i *impl) JobReq2K8sConvert(jobReq *job.CreateJobRequest) *batchv1.Job {
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobReq.Base.Name,
			Namespace: jobReq.Base.Namespace,
			Labels:    i.getk8sJobLabels(jobReq.Base.Labels),
		},
		Spec: batchv1.JobSpec{
			Completions: &jobReq.Base.Completions,
			Template:    i.getk8sJobTemplate(jobReq.Template),
		},
	}
}

// 标签转换
func (i *impl) getk8sJobLabels(jobReqLabels []*job.ListMapItem) map[string]string {
	k8sJobLabels := make(map[string]string)
	for _, label := range jobReqLabels {
		k8sJobLabels[label.Key] = label.Value
	}
	return k8sJobLabels
}

// template转换
func (i *impl) getk8sJobTemplate(jobReqTemplate *pod.Pod) corev1.PodTemplateSpec {
	k8sPod := i.podSvc.PodReq2K8s(jobReqTemplate)
	return corev1.PodTemplateSpec{
		ObjectMeta: k8sPod.ObjectMeta,
		Spec:       k8sPod.Spec,
	}
}
