package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mpaas/apps/secret"
	"go.mongodb.org/mongo-driver/bson"
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
func (i *impl) DeleteSecret(ctx context.Context, in *secret.DeleteSecretRequest) (*secret.SecretSetItem, error) {
	req := secret.NewDescribeSecretRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	secret, err := i.DescribeSecret(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] secret not found", in.Namespace, in.Name)
	}
	secretApi := i.clientSet.CoreV1().Secrets(secret.Namespace)
	err = secretApi.Delete(ctx, secret.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] secret delete fail", secret.Namespace, secret.Name)
	}
	// 从库中删除
	_, err = i.col.DeleteOne(ctx, bson.M{"name": secret.Name})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb fail", secret.Namespace, secret.Name)
	}
	return secret, nil
}

// 修改Secret
func (i *impl) UpdateSecret(ctx context.Context, in *secret.UpdateSecretRequest) (*secret.Secret, error) {
	k8sSecret := i.SecretReq2K8sConvert(in.Secret)
	secretApi := i.clientSet.CoreV1().Secrets(in.Secret.Namespace)
	if _, err := secretApi.Get(ctx, in.Secret.Name, metav1.GetOptions{}); err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] secret not found", in.Secret.Namespace, in.Secret.Name)
	}
	_, err := secretApi.Update(ctx, k8sSecret, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] secret update error, err: %s", in.Secret.Namespace, in.Secret.Name, err.Error())
	}
	secret := secret.NewDefaultSecret()
	err = i.col.FindOne(ctx, bson.M{"name": k8sSecret.Name}).Decode(secret)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sSecret.Namespace, k8sSecret.Name, err.Error())
	}
	secret.Meta.UpdatedAt = time.Now().Unix()
	secret.Secret = in.Secret
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"name": k8sSecret.Name}, bson.M{"$set": secret})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sSecret.Namespace, k8sSecret.Name, err.Error())
	}
	return secret, nil
}

// 查询Secret
func (i *impl) QuerySecret(ctx context.Context, in *secret.QuerySecretRequest) (*secret.SecretSet, error) {
	secretApi := i.clientSet.CoreV1().Secrets(in.Namespace)
	k8sSecretList, err := secretApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get secret list error, err: %s", err.Error())
	}
	secretSet := secret.NewSecretSet()
	for _, k8sConfigMap := range k8sSecretList.Items {
		// 数据转换
		secretRes := i.SecretK8s2ResItemConvert(k8sConfigMap)
		if strings.Contains(secretRes.Name, in.Keyword) {
			secretSet.AddItems(secretRes)
		}
	}
	secretSet.Total = int64(len(secretSet.Items))
	return secretSet, nil
}

// 查询Secret详情
func (i *impl) DescribeSecret(ctx context.Context, in *secret.DescribeSecretRequest) (*secret.SecretSetItem, error) {
	secretApi := i.clientSet.CoreV1().Secrets(in.Namespace)
	k8sSecret, err := secretApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get secret detail error, err: %s", err.Error())
	}
	secret := i.SecretK8s2ResDetailConvert(*k8sSecret)
	return secret, nil
}
