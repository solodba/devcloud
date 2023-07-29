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

func TestQueryRepositories(t *testing.T) {
	req := harbor.NewQueryRepositoriesRequest()
	req.CurPage = 1
	req.PageSize = 10
	req.ProjectName = "test"
	repositories, err := svc.QueryRepositories(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(repositories))
}
