package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/cronjob"
)

// 创建CronJob
func (i *impl) CreateCronJob(ctx context.Context, in *cronjob.CreateCronJobRequest) (*cronjob.CronJob, error) {
	return nil, nil
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
