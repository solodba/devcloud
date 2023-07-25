package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/job"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestCreateJob(t *testing.T) {
	req := job.NewCreateJobRequest()
	dj, err := ioutil.ReadFile("job.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	job, err := svc.CreateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(job))
}

func TestUpdateJob(t *testing.T) {
	req := job.NewUpdateJobRequest()
	dj, err := ioutil.ReadFile("job.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.Job)
	if err != nil {
		t.Fatal(err)
	}
	job, err := svc.UpdateJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(job))
}
