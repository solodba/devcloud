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
