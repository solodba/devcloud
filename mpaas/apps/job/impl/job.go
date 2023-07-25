package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/job"
)

// 创建Job
func (i *impl) CreateJob(ctx context.Context, in *job.CreateJobRequest) (*job.Job, error) {
	return nil, nil
}

// 删除Job
func (i *impl) DeleteJob(ctx context.Context, in *job.DeleteJobRequest) (*job.CreateJobRequest, error) {
	return nil, nil
}

// 修改Job
func (i *impl) UpdateJob(ctx context.Context, in *job.UpdateJobRequest) (*job.Job, error) {
	return nil, nil
}

// 查询Job
func (i *impl) QueryJob(ctx context.Context, in *job.QueryJobRequest) (*job.JobSet, error) {
	return nil, nil
}

// 查询Job详情
func (i *impl) DescribeJob(ctx context.Context, in *job.DescribeJobRequest) (*job.CreateJobRequest, error) {
	return nil, nil
}
