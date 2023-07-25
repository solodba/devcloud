package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/job"
	"github.com/solodba/devcloud/mpaas/apps/pod"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

func (i *impl) JobK8s2ReqConvert(k8sJob *batchv1.Job) *job.CreateJobRequest {
	return &job.CreateJobRequest{
		Base: &job.JobBase{
			Name:        k8sJob.Name,
			Namespace:   k8sJob.Namespace,
			Labels:      i.getJobReqLabels(k8sJob.Labels),
			Completions: i.getJobReqCompletions(k8sJob.Spec.Completions),
		},
		Template: i.getJobReqTemplate(k8sJob.Spec.Template),
	}
}

func (i *impl) getJobReqLabels(k8sJobLabels map[string]string) []*job.ListMapItem {
	jobReqLabels := make([]*job.ListMapItem, 0)
	for k, v := range k8sJobLabels {
		jobReqLabels = append(jobReqLabels, &job.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return jobReqLabels
}

func (i *impl) getJobReqCompletions(k8sJobCompletion *int32) int32 {
	var completions int32
	if k8sJobCompletion == nil {
		return completions
	}
	completions = *k8sJobCompletion
	return completions
}

func (i *impl) getJobReqTemplate(k8sJobTemplateSpec corev1.PodTemplateSpec) *pod.Pod {
	var k8sPod corev1.Pod
	k8sPod.TypeMeta.APIVersion = "v1"
	k8sPod.TypeMeta.Kind = "Pod"
	k8sPod.ObjectMeta = k8sJobTemplateSpec.ObjectMeta
	k8sPod.Spec = k8sJobTemplateSpec.Spec
	podReq := i.podSvc.PodK8s2Req(&k8sPod)
	return podReq
}
