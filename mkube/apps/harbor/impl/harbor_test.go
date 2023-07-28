package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mkube/apps/harbor"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestQueryProjects(t *testing.T) {
	req := harbor.NewQueryProjectsRequest()
	req.CurPage = 1
	req.PageSize = 2
	req.Keyword = "test"
	projects, err := svc.QueryProjects(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(projects))
}
