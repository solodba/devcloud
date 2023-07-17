package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/solodba/devcloud/mpaas/apps/pod"
	"go.mongodb.org/mongo-driver/bson"
	corev1 "k8s.io/api/core/v1"
	k8serror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

// 创建Pod
func (i *impl) CreatePod(ctx context.Context, in *pod.CreatePodRequest) (*pod.Pod, error) {
	// Pod结构体初始化
	pod, err := pod.NewPod(in)
	if err != nil {
		return nil, err
	}
	// Pod结构体转换
	k8sPod := i.PodReq2K8s(in.Pod)
	podApi := i.clientSet.CoreV1().Pods(k8sPod.Namespace)
	if getPod, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] pod already exists", getPod.Namespace, getPod.Name)
	}
	_, err = podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] create failed, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
	}
	// 新增Pod数据入库
	_, err = i.col.InsertOne(ctx, pod)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] insert into mongodb failed, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
	}
	return pod, nil
}

// 删除Pod
func (i *impl) DeletePod(ctx context.Context, in *pod.DeletePodRequest) (*pod.Pod, error) {
	req := pod.NewDescribePodRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	pod, err := i.DescribePod(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] pod not found, err: %s", in.Namespace, in.Name, err.Error())
	}
	var gracePeriodSeconds int64 = 0
	background := metav1.DeletePropagationBackground
	podApi := i.clientSet.CoreV1().Pods(pod.Base.Namespace)
	err = podApi.Delete(ctx, pod.Base.Name, metav1.DeleteOptions{
		GracePeriodSeconds: &gracePeriodSeconds,
		PropagationPolicy:  &background,
	})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete failed, err: %s", pod.Base.Namespace, pod.Base.Name, err.Error())
	}
	// 数据库中删除
	filter := bson.M{"name": pod.Base.Name}
	_, err = i.col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb failed, err: %s", pod.Base.Namespace, pod.Base.Name, err.Error())
	}
	return pod, nil
}

// 修改Pod
func (i *impl) UpdatePod(ctx context.Context, in *pod.UpdatePodRequest) (*pod.Pod, error) {
	// Pod结构体转换
	k8sPod := i.PodReq2K8s(in.Pod)
	podApi := i.clientSet.CoreV1().Pods(k8sPod.Namespace)
	getK8sPod, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] pod not exists, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
	}
	// delete+create方式更新Pod
	// pod参数校验
	k8sPodCopy := *k8sPod
	k8sPodCopy.Name = k8sPod.Name + "-validate"
	_, err = podApi.Create(ctx, &k8sPodCopy, metav1.CreateOptions{
		DryRun: []string{metav1.DryRunAll},
	})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s,name=%s] pod update failed, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
	}
	// 删除 强制删除
	background := metav1.DeletePropagationBackground
	var gracePeriodSeconds int64 = 0
	err = podApi.Delete(ctx, k8sPod.Name, metav1.DeleteOptions{
		GracePeriodSeconds: &gracePeriodSeconds,
		PropagationPolicy:  &background,
	})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s,name=%s] pod update failed, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
	}
	// Pod处于terminating状态, 监听Pod删除完毕后, 才开始创建Pod
	var labelSelector []string
	for k, v := range getK8sPod.Labels {
		labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
	}
	watcher, err := podApi.Watch(ctx, metav1.ListOptions{
		LabelSelector: strings.Join(labelSelector, ","),
	})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s,name=%s] pod update failed, detail: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
	}
	for event := range watcher.ResultChan() {
		k8sPodChan := event.Object.(*corev1.Pod)
		// 查询k8s判断是否已经删除, 那么就不用判断删除事件了
		if _, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); k8serror.IsNotFound(err) {
			//重新创建Pod
			_, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
			if err != nil {
				return nil, fmt.Errorf("[namespace=%s,name=%s] pod update failed, detail: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
			} else {
				filter := bson.M{"name": k8sPod.Name}
				_, err = i.col.UpdateOne(ctx, filter, bson.M{"": in.Pod})
				if err != nil {
					return nil, fmt.Errorf("[namespace=%s, name=%s] update from mongodb failed, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
				}
				return in.Pod, nil
			}
		}
		switch event.Type {
		case watch.Deleted:
			if k8sPodChan.Name != k8sPod.Name {
				continue
			}
			//重新创建Pod
			_, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
			if err != nil {
				return nil, fmt.Errorf("[namespace=%s,name=%s] pod update failed, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
			} else {
				filter := bson.M{"name": k8sPod.Name}
				_, err = i.col.UpdateOne(ctx, filter, bson.M{"": in.Pod})
				if err != nil {
					return nil, fmt.Errorf("[namespace=%s, name=%s] update from mongodb failed, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
				}
				return in.Pod, nil
			}
		}
	}
	return nil, nil
}

// 查询Pod
func (i *impl) QueryPod(ctx context.Context, in *pod.QueryPodRequest) (*pod.PodSet, error) {
	list, err := i.clientSet.CoreV1().Pods(in.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get pod list error, err: %s", err.Error())
	}
	podSet := pod.NewPodSet()
	for _, item := range list.Items {
		if in.NodeName != "" && item.Spec.NodeName != in.NodeName {
			continue
		}
		if strings.Contains(item.Name, in.Keyword) {
			// K8S中Pod转换成自定义Pod
			podItem := i.PodK8s2ItemRes(item)
			podSet.AddItems(podItem)
		}
	}
	podSet.Total = int64(len(podSet.PodListItem))
	return podSet, nil
}

// 查询Pod详情
func (i *impl) DescribePod(ctx context.Context, in *pod.DescribePodRequest) (*pod.Pod, error) {
	podApi := i.clientSet.CoreV1().Pods(in.Namespace)
	k8sGetPod, err := podApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s,name=%s] get pod detail error, err: %s", in.Namespace, in.Name, err.Error())
	}
	// K8s中Pod转换成自定义Pod
	return i.PodK8s2Req(k8sGetPod), nil
}
