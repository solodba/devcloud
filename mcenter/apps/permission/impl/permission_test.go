package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/namespace"
	"github.com/solodba/devcloud/mcenter/apps/permission"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestCheckPermission(t *testing.T) {
	req := permission.NewCheckPermissionRequest()
	req.UserId = "cjdfgs0eaqnmjv6mnb7g"
	req.HttpMethod = "POST"
	req.HttpPath = "/mpaas/api/v1/cluster"
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.ServiceId = "cjd1nu0eaqnlb444u060"
	permission, err := svc.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(permission))
}
