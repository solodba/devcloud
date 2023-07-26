package impl

import (
	"github.com/solodba/devcloud/mkube/apps/cronjob"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8S中结构体转换成自定义结构体
func (i *impl) CronJobK8s2ResConvert(k8sCronJob batchv1.CronJob) *cronjob.CronJobSetItem {
	return &cronjob.CronJobSetItem{
		Name:         k8sCronJob.Name,
		Namespace:    k8sCronJob.Namespace,
		Schedule:     k8sCronJob.Spec.Schedule,
		Suspend:      *k8sCronJob.Spec.Suspend,
		Active:       int64(len(k8sCronJob.Status.Active)),
		LastSchedule: i.getCronJobResLastSchedule(k8sCronJob.Status.LastScheduleTime),
		Age:          k8sCronJob.CreationTimestamp.Unix(),
	}
}

// LastSchedule转换
func (i *impl) getCronJobResLastSchedule(k8sCronJobLastSchedule *metav1.Time) int64 {
	var lastSchedule int64
	if k8sCronJobLastSchedule == nil {
		return lastSchedule
	}
	lastSchedule = (*k8sCronJobLastSchedule).Unix()
	return lastSchedule
}
