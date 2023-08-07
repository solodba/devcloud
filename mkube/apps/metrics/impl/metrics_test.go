package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestQueryClusterUsage(t *testing.T) {
	req := metrics.NewQueryClusterUsageRequest()
	metricSet, err := svc.QueryClusterUsage(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(metricSet))
}

func TestQueryResource(t *testing.T) {
	req := metrics.NewQueryResourceRequest()
	metricSet, err := svc.QueryResource(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(metricSet))
}

func TestQueryClusterInfo(t *testing.T) {
	req := metrics.NewQueryClusterInfoRequest()
	metricSet, err := svc.QueryClusterInfo(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(metricSet))
}
