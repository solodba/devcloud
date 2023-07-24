package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mpaas/apps/configmap"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建ConfigMap
func (i *impl) CreateConfigMap(ctx context.Context, in *configmap.CreateConfigMapRequest) (*configmap.ConfigMap, error) {
	// ConfigMap转换
	k8sConfigMap := i.CmReq2K8s(in)
	configmapApi := i.clientSet.CoreV1().ConfigMaps(k8sConfigMap.Namespace)
	// 判断ConfigMap是否存在
	if getK8sConfigMap, err := configmapApi.Get(ctx, k8sConfigMap.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap already exists", getK8sConfigMap.Namespace, getK8sConfigMap.Name)
	}
	// 创建ConfigMap
	_, err := configmapApi.Create(ctx, k8sConfigMap, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap create fail", k8sConfigMap.Namespace, k8sConfigMap.Name)
	}
	configmap := configmap.NewConfigMap(in)
	// 入库
	_, err = i.col.InsertOne(ctx, configmap)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap insert mongodb fail, err: %s", k8sConfigMap.Namespace, k8sConfigMap.Name, err.Error())
	}
	return configmap, nil
}

// 删除ConfigMap
func (i *impl) DeleteConfigMap(ctx context.Context, in *configmap.DeleteConfigMapRequest) (*configmap.ConfigMapSetItem, error) {
	req := configmap.NewDescribeConfigMapRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	configmap, err := i.DescribeConfigMap(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap not found", in.Namespace, in.Name)
	}
	configmapAPi := i.clientSet.CoreV1().ConfigMaps(configmap.Namespace)
	err = configmapAPi.Delete(ctx, configmap.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap delete fail", configmap.Namespace, configmap.Name)
	}
	// 从库中删除
	_, err = i.col.DeleteOne(ctx, bson.M{"name": configmap.Name})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb fail", configmap.Namespace, configmap.Name)
	}
	return configmap, nil
}

// 更新ConfigMap
func (i *impl) UpdateConfigMap(ctx context.Context, in *configmap.UpdateConfigMapRequest) (*configmap.ConfigMap, error) {
	k8sConfigMap := i.CmReq2K8s(in.ConfigMap)
	configmapApi := i.clientSet.CoreV1().ConfigMaps(in.ConfigMap.Namespace)
	if _, err := configmapApi.Get(ctx, in.ConfigMap.Name, metav1.GetOptions{}); err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap not found", in.ConfigMap.Namespace, in.ConfigMap.Name)
	}
	_, err := configmapApi.Update(ctx, k8sConfigMap, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] configmap update error, err: %s", in.ConfigMap.Namespace, in.ConfigMap.Name, err.Error())
	}
	configmap := configmap.NewDefaultConfigMap()
	err = i.col.FindOne(ctx, bson.M{"name": k8sConfigMap.Name}).Decode(configmap)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sConfigMap.Namespace, k8sConfigMap.Name, err.Error())
	}
	configmap.Meta.UpdatedAt = time.Now().Unix()
	configmap.ConfigMap = in.ConfigMap
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"name": k8sConfigMap.Name}, bson.M{"$set": configmap})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sConfigMap.Namespace, k8sConfigMap.Name, err.Error())
	}
	return configmap, nil
}

// 查询ConfigMap
func (i *impl) QueryConfigMap(ctx context.Context, in *configmap.QueryConfigMapRequest) (*configmap.ConfigMapSet, error) {
	configmapApi := i.clientSet.CoreV1().ConfigMaps(in.Namespace)
	k8sConfigMapList, err := configmapApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get configmap list error, err: %s", err.Error())
	}
	configmapSet := configmap.NewConfigMapSet()
	for _, k8sConfigMap := range k8sConfigMapList.Items {
		// 数据转换
		configMapRes := i.GetCmResItem(&k8sConfigMap)
		if strings.Contains(configMapRes.Name, in.Keyword) {
			configmapSet.AddItems(configMapRes)
		}
	}
	configmapSet.Total = int64(len(configmapSet.Items))
	return configmapSet, nil
}

// 查询ConfigMap详情
func (i *impl) DescribeConfigMap(ctx context.Context, in *configmap.DescribeConfigMapRequest) (*configmap.ConfigMapSetItem, error) {
	configmapApi := i.clientSet.CoreV1().ConfigMaps(in.Namespace)
	k8sConfigMap, err := configmapApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get configmap detail error, err: %s", err.Error())
	}
	configmap := i.GetCmResDetail(k8sConfigMap)
	return configmap, nil
}
