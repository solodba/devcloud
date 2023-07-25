package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/cronjob"
	"github.com/solodba/devcloud/mpaas/apps/pod"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

// K8S中CronJob结构体转换成自定义结构体
func (i *impl) CronJobK8s2ReqConvert(k8sCronJob *batchv1.CronJob) *cronjob.CreateCronJobRequest {
	return &cronjob.CreateCronJobRequest{
		Base: &cronjob.CronJobBase{
			Name:                       k8sCronJob.Name,
			Namespace:                  k8sCronJob.Namespace,
			Labels:                     i.getCronJobReqLabels(k8sCronJob.Labels),
			Schedule:                   k8sCronJob.Spec.Schedule,
			Suspend:                    i.getCronJobReqSuspend(k8sCronJob.Spec.Suspend),
			ConcurrencyPolicy:          string(k8sCronJob.Spec.ConcurrencyPolicy),
			SuccessfulJobsHistoryLimit: i.getCronJobReqSuccessfulJobsHistoryLimit(k8sCronJob.Spec.SuccessfulJobsHistoryLimit),
			FailedJobsHistoryLimit:     i.getCronJobReqFailedJobsHistoryLimit(k8sCronJob.Spec.FailedJobsHistoryLimit),
			JobBase: &cronjob.JobBase{
				Completions:  i.getCronJobReqCompletions(k8sCronJob.Spec.JobTemplate.Spec.Completions),
				BackoffLimit: i.getCronJobReqBackoffLimit(k8sCronJob.Spec.JobTemplate.Spec.BackoffLimit),
			},
		},
		Template: i.getCronJobReqTemplate(k8sCronJob.Spec.JobTemplate.Spec.Template),
	}
}

// 标签转换
func (i *impl) getCronJobReqLabels(k8sCronJobLabels map[string]string) []*cronjob.ListMapItem {
	cronJobReqLabels := make([]*cronjob.ListMapItem, 0)
	for k, v := range k8sCronJobLabels {
		cronJobReqLabels = append(cronJobReqLabels, &cronjob.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return cronJobReqLabels
}

// Suspend转换
func (i *impl) getCronJobReqSuspend(k8sCronJobSuspend *bool) bool {
	var suspend bool
	if k8sCronJobSuspend == nil {
		return suspend
	}
	suspend = *k8sCronJobSuspend
	return suspend
}

// JobHistoryLimit转换
func (i *impl) getCronJobReqSuccessfulJobsHistoryLimit(k8sCronJobSuccess *int32) int32 {
	var success int32
	if k8sCronJobSuccess == nil {
		return success
	}
	success = *k8sCronJobSuccess
	return success
}

// FailedJobsHistoryLimit转换
func (i *impl) getCronJobReqFailedJobsHistoryLimit(k8sCronJobFailed *int32) int32 {
	var failed int32
	if k8sCronJobFailed == nil {
		return failed
	}
	failed = *k8sCronJobFailed
	return failed
}

// Completions转换
func (i *impl) getCronJobReqCompletions(k8sCronJobCompletions *int32) int32 {
	var completions int32
	if k8sCronJobCompletions == nil {
		return completions
	}
	completions = *k8sCronJobCompletions
	return completions
}

// BackoffLimit转换
func (i *impl) getCronJobReqBackoffLimit(k8sCronJobBackoffLimit *int32) int32 {
	var backOffLimit int32
	if k8sCronJobBackoffLimit == nil {
		return backOffLimit
	}
	backOffLimit = *k8sCronJobBackoffLimit
	return backOffLimit
}

// Template转换
func (i *impl) getCronJobReqTemplate(k8sCronJobTemplate corev1.PodTemplateSpec) *pod.Pod {
	var k8sPod corev1.Pod
	k8sPod.TypeMeta.APIVersion = "v1"
	k8sPod.TypeMeta.Kind = "Pod"
	k8sPod.ObjectMeta = k8sCronJobTemplate.ObjectMeta
	k8sPod.Spec = k8sCronJobTemplate.Spec
	podReq := i.podSvc.PodK8s2Req(&k8sPod)
	return podReq
}
