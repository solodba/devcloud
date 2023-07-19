package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/configmap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建ConfigMap
func (i *impl) CreateConfigMap(ctx context.Context, in *configmap.CreateConfigMapRequest) (*configmap.ConfigMap, error) {
	// ConfigMap转换
	k8sConfigMap := i.CmReq2K8s(in)
	svcApi := i.clientSet.CoreV1().ConfigMaps(k8sConfigMap.Namespace)
	// 判断Service是否存在
	if getSVC, err := svcApi.Get(ctx, k8sConfigMap.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap already exists", getSVC.Namespace, getSVC.Name)
	}
	// 创建Service
	_, err := svcApi.Create(ctx, k8sConfigMap, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap create fail", k8sConfigMap.Namespace, k8sConfigMap.Name)
	}
	configmap := configmap.NewConfigMap(in)
	// 入库
	_, err = i.col.InsertOne(ctx, configmap)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] service insert mongodb fail, err: %s", k8sConfigMap.Namespace, k8sConfigMap.Name, err.Error())
	}
	return configmap, nil
}

// 删除ConfigMap
func (i *impl) DeleteConfigMap(ctx context.Context, in *configmap.DeleteConfigMapRequest) (*configmap.ConfigMap, error) {
	return nil, nil
}

// 更新ConfigMap
func (i *impl) UpdateConfigMap(ctx context.Context, in *configmap.UpdateConfigMapRequest) (*configmap.ConfigMap, error) {
	return nil, nil
}

// 查询ConfigMap
func (i *impl) QueryConfigMap(ctx context.Context, in *configmap.QueryConfigMapRequest) (*configmap.ConfigMapSet, error) {
	return nil, nil
}

// 查询ConfigMap详情
func (i *impl) DescribeConfigMap(ctx context.Context, in *configmap.DescribeConfigMapRequest) (*configmap.ConfigMap, error) {
	return nil, nil
}
