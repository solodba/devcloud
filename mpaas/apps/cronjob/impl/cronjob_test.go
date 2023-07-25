package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/cronjob"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestCreateCronJob(t *testing.T) {
	req := cronjob.NewCreateCronJobRequest()
	dj, err := ioutil.ReadFile("cronjob.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	cronjob, err := svc.CreateCronJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(cronjob))
}

func TestUpdateCronJob(t *testing.T) {
	req := cronjob.NewUpdateCronJobRequest()
	dj, err := ioutil.ReadFile("cronjob.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.CronJob)
	if err != nil {
		t.Fatal(err)
	}
	cronjob, err := svc.UpdateCronJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(cronjob))
}

func TestQueryCronJob(t *testing.T) {
	req := cronjob.NewQueryCronJobRequest()
	req.Namespace = "test"
	req.Keyword = "nginx"
	cronjobSet, err := svc.QueryCronJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(cronjobSet))
}

func TestDescribeJob(t *testing.T) {
	req := cronjob.NewDescribeCronJobRequest()
	req.Namespace = "test"
	req.Name = "nginx-cronjob"
	cronjob, err := svc.DescribeCronJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(cronjob))
}

func TestDeleteCronJob(t *testing.T) {
	req := cronjob.NewDeleteCronJobRequest()
	req.Namespace = "test"
	req.Name = "nginx-cronjob"
	cronjob, err := svc.DeleteCronJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(cronjob))
}
