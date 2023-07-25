package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/cronjob"
	"github.com/solodba/devcloud/mpaas/apps/pod"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义CronJob结构体转换成K8S中结构体
func (i *impl) CronJobReq2K8sConvert(cronjobReq *cronjob.CreateCronJobRequest) *batchv1.CronJob {
	return &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cronjobReq.Base.Name,
			Namespace: cronjobReq.Base.Namespace,
			Labels:    i.getK8sCronJobLabels(cronjobReq.Base.Labels),
		},
		Spec: batchv1.CronJobSpec{
			Schedule:                   cronjobReq.Base.Schedule,
			Suspend:                    &cronjobReq.Base.Suspend,
			ConcurrencyPolicy:          batchv1.ConcurrencyPolicy(cronjobReq.Base.ConcurrencyPolicy),
			SuccessfulJobsHistoryLimit: &cronjobReq.Base.SuccessfulJobsHistoryLimit,
			FailedJobsHistoryLimit:     &cronjobReq.Base.FailedJobsHistoryLimit,
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					// Selector:     pc.getK8sCronJobSelector(cronjobReq.Base.Selector),
					Completions:  &cronjobReq.Base.JobBase.Completions,
					BackoffLimit: &cronjobReq.Base.JobBase.BackoffLimit,
					Template:     i.getk8sCronJobTemplate(cronjobReq.Template),
				},
			},
		},
	}
}

// 标签转换
func (i *impl) getK8sCronJobLabels(cronJobReqLabels []*cronjob.ListMapItem) map[string]string {
	k8sCronJobLabels := make(map[string]string)
	for _, label := range cronJobReqLabels {
		k8sCronJobLabels[label.Key] = label.Value
	}
	return k8sCronJobLabels
}

// 标签选择器转换
func (i *impl) getK8sCronJobSelector(cronJobReqSelector []*cronjob.ListMapItem) *metav1.LabelSelector {
	var labelSelector metav1.LabelSelector
	labelSelector.MatchLabels = make(map[string]string)
	for _, selector := range cronJobReqSelector {
		labelSelector.MatchLabels[selector.Key] = selector.Value
	}
	return &labelSelector
}

// Template转换
func (i *impl) getk8sCronJobTemplate(cronJobReqTemplate *pod.Pod) corev1.PodTemplateSpec {
	k8sPod := i.podSvc.PodReq2K8s(cronJobReqTemplate)
	return corev1.PodTemplateSpec{
		ObjectMeta: k8sPod.ObjectMeta,
		Spec:       k8sPod.Spec,
	}
}
