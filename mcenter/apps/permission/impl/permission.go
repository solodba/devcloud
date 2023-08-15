package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/permission"
)

// 用户访问鉴权
func (i *impl) CheckPermission(ctx context.Context, in *permission.CheckPermissionRequest) (*permission.CheckPermissionResponse, error) {
	return nil, nil
}
