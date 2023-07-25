package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/job"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *impl) JobK8s2ResConvert(k8sJob batchv1.Job) *job.JobSetItem {
	return &job.JobSetItem{
		Name:        k8sJob.Name,
		Namespace:   k8sJob.Namespace,
		Completions: i.getJobResCompletions(k8sJob.Spec.Completions),
		Succeeded:   k8sJob.Status.Succeeded,
		Age:         k8sJob.CreationTimestamp.Unix(),
		Duration:    i.getJobResCompletionTime(k8sJob.Status.CompletionTime) - i.getJobResStartTime(k8sJob.Status.StartTime),
	}
}

func (i *impl) getJobResCompletions(k8sJobCompletions *int32) int32 {
	var completions int32
	if k8sJobCompletions == nil {
		return completions
	}
	completions = *k8sJobCompletions
	return completions
}

func (i *impl) getJobResCompletionTime(completionTime *metav1.Time) int64 {
	var ctime int64
	if completionTime == nil {
		return ctime
	}
	ctime = (*completionTime).Unix()
	return ctime
}

func (i *impl) getJobResStartTime(startTime *metav1.Time) int64 {
	var stime int64
	if startTime == nil {
		return stime
	}
	stime = (*startTime).Unix()
	return stime
}
