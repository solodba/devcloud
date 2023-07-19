package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/configmap"
)

// 创建ConfigMap
func (i *impl) CreateConfigMap(ctx context.Context, in *configmap.CreateConfigMapRequest) (*configmap.ConfigMap, error) {
	return nil, nil
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
