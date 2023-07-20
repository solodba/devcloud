package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/secret"
)

// 创建Secret
func (i *impl) CreateSecret(ctx context.Context, in *secret.CreateSecretRequest) (*secret.Secret, error) {
	return nil, nil
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
