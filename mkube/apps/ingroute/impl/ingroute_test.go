package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mkube/apps/ingroute"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestDescribeIngressRoute(t *testing.T) {
	req := ingroute.NewDescribeIngressRouteRequest()
	req.Namespace = "test"
	req.Name = "web-ingressroute"
	ingressRoute, err := svc.DescribeIngressRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingressRoute))
}
