package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/secret"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建Secret
func (i *impl) CreateSecret(ctx context.Context, in *secret.CreateSecretRequest) (*secret.Secret, error) {
	// Secret转换
	k8sSecret := i.SecretReq2K8sConvert(in)
	secretApi := i.clientSet.CoreV1().Secrets(k8sSecret.Namespace)
	// 判断Secret是否存在
	if getSecret, err := secretApi.Get(ctx, k8sSecret.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] secret already exists", getSecret.Namespace, getSecret.Name)
	}
	// 创建Secret
	_, err := secretApi.Create(ctx, k8sSecret, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] secret create fail", k8sSecret.Namespace, k8sSecret.Name)
	}
	secret := secret.NewSecret(in)
	// 入库
	_, err = i.col.InsertOne(ctx, secret)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] secret insert mongodb fail, err: %s", k8sSecret.Namespace, k8sSecret.Name, err.Error())
	}
	return secret, nil
}

// 删除Secret
func (i *impl) DeleteSecret(ctx context.Context, in *secret.DeleteSecretRequest) (*secret.Secret, error) {
	return nil, nil
}

// 修改Secret
func (i *impl) UpdateSecret(ctx context.Context, in *secret.UpdateSecretRequest) (*secret.Secret, error) {
	return nil, nil
}

// 查询Secret
func (i *impl) QuerySecret(ctx context.Context, in *secret.QuerySecretRequest) (*secret.SecretSetItem, error) {
	return nil, nil
}

// 查询Secret详情
func (i *impl) DescribeSecret(ctx context.Context, in *secret.DescribeSecretRequest) (*secret.SecretSet, error) {
	return nil, nil
}
