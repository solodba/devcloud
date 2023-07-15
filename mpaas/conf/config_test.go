package conf_test

import (
	"testing"

	"github.com/solodba/devcloud/tree/main/mpaas/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().App.Name)
	t.Log(conf.C().App.Http)
	t.Log(conf.C().App.Grpc)
	t.Log(conf.C().MongoDB)
	t.Log(conf.C().MongoDB.GetDB())
	t.Log(conf.C().K8s)
	t.Log(conf.C().K8s.GetK8sConn())
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().App.Name)
	t.Log(conf.C().App.Http)
	t.Log(conf.C().App.Grpc)
	t.Log(conf.C().MongoDB)
	t.Log(conf.C().MongoDB.GetDB())
	t.Log(conf.C().K8s)
	t.Log(conf.C().K8s.GetK8sConn())
}
