package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mkube/apps/prometheus"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestQueryMetricsFromProm(t *testing.T) {
	req := prometheus.NewQueryMetricsFromPromRequest("cluster_mem")
	promeMetrics, err := svc.QueryMetricsFromProm(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(promeMetrics))
}

func TestQueryClusterUsageRange(t *testing.T) {
	req := prometheus.NewQueryClusterUsageRangeRequest()
	metrics, err := svc.QueryClusterUsageRange(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(metrics))
}
