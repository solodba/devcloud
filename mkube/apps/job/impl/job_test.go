package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mkube/apps/job"
	"github.com/solodba/devcloud/mkube/test/tools"
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

func TestQuertJob(t *testing.T) {
	req := job.NewQueryJobRequest()
	req.Namespace = "test"
	req.Keyword = "nginx-job"
	jobSet, err := svc.QueryJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(jobSet))
}

func TestDescribeJob(t *testing.T) {
	req := job.NewDescribeJobRequest()
	req.Namespace = "test"
	req.Name = "nginx-job"
	job, err := svc.DescribeJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(job))
}

func TestDeleteJob(t *testing.T) {
	req := job.NewDeleteJobRequest()
	req.Namespace = "test"
	req.Name = "nginx-job"
	job, err := svc.DeleteJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(job))
}
