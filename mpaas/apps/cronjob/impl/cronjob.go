package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/cronjob"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建CronJob
func (i *impl) CreateCronJob(ctx context.Context, in *cronjob.CreateCronJobRequest) (*cronjob.CronJob, error) {
	// Job转换
	k8sCronJob := i.CronJobReq2K8sConvert(in)
	cronJobApi := i.clientSet.BatchV1().CronJobs(k8sCronJob.Namespace)
	// 判断Job是否存在
	if getK8sCronJob, err := cronJobApi.Get(ctx, k8sCronJob.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] cronjob already exists", getK8sCronJob.Namespace, getK8sCronJob.Name)
	}
	// 创建Job
	_, err := cronJobApi.Create(ctx, k8sCronJob, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] cronjob create fail", k8sCronJob.Namespace, k8sCronJob.Name)
	}
	cronjob := cronjob.NewCronJob(in)
	// 入库
	_, err = i.col.InsertOne(ctx, cronjob)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] cronjob insert mongodb fail, err: %s", k8sCronJob.Namespace, k8sCronJob.Name, err.Error())
	}
	return cronjob, nil
}

// 删除CronJob
func (i *impl) DeleteCronJob(ctx context.Context, in *cronjob.DeleteCronJobRequest) (*cronjob.CronJob, error) {
	return nil, nil
}

// 修改CronJob
func (i *impl) UpdateCronJob(ctx context.Context, in *cronjob.UpdateCronJobRequest) (*cronjob.CronJob, error) {
	return nil, nil
}

// 查询CronJob
func (i *impl) QueryCronJob(ctx context.Context, in *cronjob.QueryCronJobRequest) (*cronjob.CronJobSet, error) {
	return nil, nil
}

// 查询CronJob详情
func (i *impl) DescribeCronJob(ctx context.Context, in *cronjob.DescribeCronJobRequest) (*cronjob.CronJob, error) {
	return nil, nil
}
