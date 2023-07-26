package impl

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mkube/apps/job"
	"go.mongodb.org/mongo-driver/bson"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/utils/pointer"
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
	req := job.NewDescribeJobRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	job, err := i.DescribeJob(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] job not found", in.Namespace, in.Name)
	}
	jobApi := i.clientSet.BatchV1().Jobs(job.Base.Namespace)
	jobK8s, err := jobApi.Get(ctx, job.Base.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	//开启监听
	var labelSelector []string
	for k, v := range jobK8s.Labels {
		labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
	}
	watcher, err := jobApi.Watch(ctx, metav1.ListOptions{
		LabelSelector: strings.Join(labelSelector, ","),
	})
	if err != nil {
		return nil, err
	}
	var podLabelSelector []string
	for k, v := range jobK8s.Spec.Template.Labels {
		podLabelSelector = append(podLabelSelector, fmt.Sprintf("%s=%s", k, v))
	}
	err = jobApi.Delete(ctx, job.Base.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}
	notify := make(chan error)
	go func(thisJob *batchv1.Job, notify chan error) {
		//监听删除信号后，创建
		for {
			select {
			case e := <-watcher.ResultChan():
				switch e.Type {
				case watch.Deleted:
					//删除关联Pod
					if list, errx := i.clientSet.CoreV1().Pods(jobK8s.Namespace).
						List(ctx, metav1.ListOptions{
							LabelSelector: strings.Join(podLabelSelector, ","),
						}); errx == nil {
						//清理job关联的Pod
						for _, item := range list.Items {
							//delete pod
							background := metav1.DeletePropagationBackground
							err = i.clientSet.CoreV1().Pods(item.Namespace).Delete(ctx, item.Name, metav1.DeleteOptions{
								GracePeriodSeconds: pointer.Int64Ptr(0),
								PropagationPolicy:  &background,
							})
						}
					}
					notify <- nil
					return
				}
			case <-time.After(5 * time.Second):
				notify <- errors.New("删除Job超时")
				return
			}
		}
	}(jobK8s, notify)
	select {
	case errx := <-notify:
		if errx != nil {
			return nil, errx
		}
	}
	// 从库中删除
	_, err = i.col.DeleteOne(ctx, bson.M{"base.name": job.Base.Name})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb fail", job.Base.Namespace, job.Base.Name)
	}
	return job, nil
}

// 修改Job
func (i *impl) UpdateJob(ctx context.Context, in *job.UpdateJobRequest) (*job.Job, error) {
	k8sJob := i.JobReq2K8sConvert(in.Job)
	jobApi := i.clientSet.BatchV1().Jobs(in.Job.Base.Namespace)
	getK8sJob, err := jobApi.Get(ctx, in.Job.Base.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] job not found", in.Job.Base.Namespace, in.Job.Base.Name)
	}
	// 更新Job
	//参数校验
	jobCopy := *k8sJob
	vJobName := k8sJob.Name + "-validate"
	jobCopy.Name = vJobName
	_, err = jobApi.Create(ctx, &jobCopy, metav1.CreateOptions{
		DryRun: []string{metav1.DryRunAll},
	})
	if err != nil {
		return nil, err
	}
	//开启监听
	var labelSelector []string
	for k, v := range getK8sJob.Labels {
		labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
	}
	var podLabelSelector []string
	for k, v := range getK8sJob.Spec.Template.Labels {
		podLabelSelector = append(podLabelSelector, fmt.Sprintf("%s=%s", k, v))
	}
	watcher, errin := jobApi.Watch(ctx, metav1.ListOptions{
		LabelSelector: strings.Join(labelSelector, ","),
	})
	if errin != nil {
		return nil, errin
	}
	notify := make(chan error)
	go func(thisJob *batchv1.Job, notify chan error) {
		//监听删除信号后，创建
		for {
			select {
			case e := <-watcher.ResultChan():
				switch e.Type {
				case watch.Deleted:
					//删除关联Pod
					if list, errx := i.clientSet.CoreV1().Pods(getK8sJob.Namespace).
						List(ctx, metav1.ListOptions{
							LabelSelector: strings.Join(podLabelSelector, ","),
						}); errx == nil {
						for _, item := range list.Items {
							//delete pod
							background := metav1.DeletePropagationBackground
							err = i.clientSet.CoreV1().Pods(item.Namespace).Delete(ctx, item.Name, metav1.DeleteOptions{
								GracePeriodSeconds: pointer.Int64Ptr(0),
								PropagationPolicy:  &background,
							})
						}
					}
					_, errin = jobApi.Create(ctx, thisJob, metav1.CreateOptions{})
					notify <- errin
					return
				}
			case <-time.After(5 * time.Second):
				notify <- errors.New("更新Job超时")
				return
				//fmt.Println("timeout")
			}
		}
	}(k8sJob, notify)
	//删除
	background := metav1.DeletePropagationForeground
	err = jobApi.Delete(ctx, k8sJob.Name, metav1.DeleteOptions{
		PropagationPolicy: &background,
	})
	if err != nil {
		return nil, err
	}
	//监听删除后重新创建的信息
	select {
	case errx := <-notify:
		if errx != nil {
			return nil, errx
		}
	}
	job := job.NewDefaultJob()
	err = i.col.FindOne(ctx, bson.M{"base.name": k8sJob.Name}).Decode(job)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sJob.Namespace, k8sJob.Name, err.Error())
	}
	job.Meta.UpdatedAt = time.Now().Unix()
	job.Job = in.Job
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"base.name": k8sJob.Name}, bson.M{"$set": job})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sJob.Namespace, k8sJob.Name, err.Error())
	}
	return job, nil
}

// 查询Job
func (i *impl) QueryJob(ctx context.Context, in *job.QueryJobRequest) (*job.JobSet, error) {
	jobApi := i.clientSet.BatchV1().Jobs(in.Namespace)
	k8sJobList, err := jobApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get job list error, err: %s", err.Error())
	}
	jobSet := job.NewJobSet()
	for _, k8sJob := range k8sJobList.Items {
		// 数据转换
		jobRes := i.JobK8s2ResConvert(k8sJob)
		if strings.Contains(jobRes.Name, in.Keyword) {
			jobSet.AddItems(jobRes)
		}
	}
	jobSet.Total = int64(len(jobSet.Items))
	return jobSet, nil
}

// 查询Job详情
func (i *impl) DescribeJob(ctx context.Context, in *job.DescribeJobRequest) (*job.CreateJobRequest, error) {
	jobApi := i.clientSet.BatchV1().Jobs(in.Namespace)
	k8sJob, err := jobApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get job detail error, err: %s", err.Error())
	}
	jobReq := i.JobK8s2ReqConvert(k8sJob)
	jobReqLables := make([]*job.ListMapItem, 0)
	for _, label := range jobReq.Base.Labels {
		if strings.Contains(label.Key, "controller-uid") {
			continue
		}
		if strings.Contains(label.Key, "job-name") {
			continue
		}
		jobReqLables = append(jobReqLables, label)
	}
	jobReq.Base.Labels = jobReqLables
	return jobReq, nil
}
