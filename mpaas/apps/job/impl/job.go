package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/job"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建Job
func (i *impl) CreateJob(ctx context.Context, in *job.CreateJobRequest) (*job.Job, error) {
	// Job转换
	k8sJob := i.JobReq2K8sConvert(in)
	jobApi := i.clientSet.BatchV1().Jobs(k8sJob.Namespace)
	// 判断Job是否存在
	if getK8sJob, err := jobApi.Get(ctx, k8sJob.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] job already exists", getK8sJob.Namespace, getK8sJob.Name)
	}
	// 创建Job
	_, err := jobApi.Create(ctx, k8sJob, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] job create fail", k8sJob.Namespace, k8sJob.Name)
	}
	job := job.NewJob(in)
	// 入库
	_, err = i.col.InsertOne(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] job insert mongodb fail, err: %s", k8sJob.Namespace, k8sJob.Name, err.Error())
	}
	return job, nil
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
