package impl

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mpaas/apps/cronjob"
	"go.mongodb.org/mongo-driver/bson"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
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
func (i *impl) DeleteCronJob(ctx context.Context, in *cronjob.DeleteCronJobRequest) (*cronjob.CreateCronJobRequest, error) {
	req := cronjob.NewDescribeCronJobRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	cronjob, err := i.DescribeCronJob(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] cronjob not found", in.Namespace, in.Name)
	}
	cronJobApi := i.clientSet.BatchV1().CronJobs(cronjob.Base.Namespace)
	cronJobK8s, err := cronJobApi.Get(ctx, cronjob.Base.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	//开启监听
	var labelSelector []string
	for k, v := range cronJobK8s.Labels {
		labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
	}
	watcher, err := cronJobApi.Watch(ctx, metav1.ListOptions{
		LabelSelector: strings.Join(labelSelector, ","),
	})
	if err != nil {
		return nil, err
	}
	err = cronJobApi.Delete(ctx, cronjob.Base.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}
	notify := make(chan error)
	go func(thisJob *batchv1.CronJob, notify chan error) {
		//监听删除信号后，创建
		for {
			select {
			case e := <-watcher.ResultChan():
				switch e.Type {
				case watch.Deleted:
					notify <- nil
					return
				}
			case <-time.After(5 * time.Second):
				notify <- errors.New("删除CronJob超时")
				return
			}
		}
	}(cronJobK8s, notify)
	select {
	case errx := <-notify:
		if errx != nil {
			return nil, errx
		}
	}
	// 从库中删除
	_, err = i.col.DeleteOne(ctx, bson.M{"base.name": cronjob.Base.Name})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb fail", cronjob.Base.Namespace, cronjob.Base.Name)
	}
	return cronjob, nil
}

// 修改CronJob
func (i *impl) UpdateCronJob(ctx context.Context, in *cronjob.UpdateCronJobRequest) (*cronjob.CronJob, error) {
	k8sCronJob := i.CronJobReq2K8sConvert(in.CronJob)
	cronJobApi := i.clientSet.BatchV1().CronJobs(in.CronJob.Base.Namespace)
	getK8sCronJob, err := cronJobApi.Get(ctx, in.CronJob.Base.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] cronjob not found", in.CronJob.Base.Namespace, in.CronJob.Base.Name)
	}
	// 更新
	//参数校验
	cronJobCopy := *k8sCronJob
	vCronJobName := cronJobCopy.Name + "-validate"
	cronJobCopy.Name = vCronJobName
	_, err = cronJobApi.Create(ctx, &cronJobCopy, metav1.CreateOptions{
		DryRun: []string{metav1.DryRunAll},
	})
	if err != nil {
		return nil, err
	}
	//开启监听
	var labelSelector []string
	for k, v := range getK8sCronJob.Labels {
		labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
	}
	var podLabelSelector []string
	for k, v := range getK8sCronJob.Spec.JobTemplate.Spec.Template.Labels {
		podLabelSelector = append(podLabelSelector, fmt.Sprintf("%s=%s", k, v))
	}
	//监听cronjob删除状态
	watcher, errin := cronJobApi.Watch(ctx, metav1.ListOptions{
		LabelSelector: strings.Join(labelSelector, ","),
	})
	if errin != nil {
		return nil, errin
	}
	notify := make(chan error)
	go func(thisCronJob *batchv1.CronJob, notify chan error) {
		//监听删除信号后，创建
		for {
			select {
			case e := <-watcher.ResultChan():
				switch e.Type {
				case watch.Deleted:
					//删除关联Job
					_, errin = cronJobApi.Create(ctx, thisCronJob, metav1.CreateOptions{})
					notify <- errin
					return
				}
			case <-time.After(5 * time.Second):
				notify <- errors.New("更新Job超时")
				return
				//fmt.Println("timeout")
			}
		}
	}(k8sCronJob, notify)
	//删除
	background := metav1.DeletePropagationForeground
	err = cronJobApi.Delete(ctx, k8sCronJob.Name, metav1.DeleteOptions{
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
	cronjob := cronjob.NewDefaultCronJob()
	err = i.col.FindOne(ctx, bson.M{"base.name": k8sCronJob.Name}).Decode(cronjob)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sCronJob.Namespace, k8sCronJob.Name, err.Error())
	}
	cronjob.Meta.UpdatedAt = time.Now().Unix()
	cronjob.CronJob = in.CronJob
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"base.name": k8sCronJob.Name}, bson.M{"$set": cronjob})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sCronJob.Namespace, k8sCronJob.Name, err.Error())
	}
	return cronjob, nil
}

// 查询CronJob
func (i *impl) QueryCronJob(ctx context.Context, in *cronjob.QueryCronJobRequest) (*cronjob.CronJobSet, error) {
	cronJobApi := i.clientSet.BatchV1().CronJobs(in.Namespace)
	k8sCronJobList, err := cronJobApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get cronjob list error, err: %s", err.Error())
	}
	cronJobSet := cronjob.NewCronJobSet()
	for _, k8sCronJob := range k8sCronJobList.Items {
		// 数据转换
		cronJobRes := i.CronJobK8s2ResConvert(k8sCronJob)
		if strings.Contains(cronJobRes.Name, in.Keyword) {
			cronJobSet.AddItems(cronJobRes)
		}
	}
	cronJobSet.Total = int64(len(cronJobSet.Items))
	return cronJobSet, nil
}

// 查询CronJob详情
func (i *impl) DescribeCronJob(ctx context.Context, in *cronjob.DescribeCronJobRequest) (*cronjob.CreateCronJobRequest, error) {
	cronJobApi := i.clientSet.BatchV1().CronJobs(in.Namespace)
	k8sCronJob, err := cronJobApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get cronjob detail error, err: %s", err.Error())
	}
	return i.CronJobK8s2ReqConvert(k8sCronJob), nil
}
