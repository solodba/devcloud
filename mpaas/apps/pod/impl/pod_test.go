package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/pod"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestCreatePod(t *testing.T) {
	req := pod.NewCreatePodRequest()
	data, err := ioutil.ReadFile("pod_createorupdate_volume.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(data, req.Pod)
	if err != nil {
		t.Fatal(err)
	}
	pod, err := svc.CreatePod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pod))
}

func TestQueryPod(t *testing.T) {
	req := pod.NewQueryPodRequest()
	req.Namespace = "test"
	req.NodeName = "node1"
	req.Keyword = "te"
	podSet, err := svc.QueryPod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(podSet))
}

func TestDescribePod(t *testing.T) {
	req := pod.NewDescribePodRequest()
	req.Name = "test"
	req.Namespace = "test"
	pod, err := svc.DescribePod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pod))
}

func TestDeletePod(t *testing.T) {
	req := pod.NewDeletePodRequest()
	req.Namespace = "test"
	req.Name = "test"
	pod, err := svc.DeletePod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pod))
}

func TestUpdatePod(t *testing.T) {
	createPodReq := pod.NewCreatePodRequest()
	data, err := ioutil.ReadFile("pod_createorupdate.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(data, createPodReq.Pod)
	if err != nil {
		t.Fatal(err)
	}
	updatePodReq := pod.NewUpdatePodRequest(createPodReq.Pod)
	pod, err := svc.UpdatePod(ctx, updatePodReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pod))
}
