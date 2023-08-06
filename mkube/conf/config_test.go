package conf_test

import (
	"testing"

	"github.com/solodba/devcloud/mkube/conf"
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
	t.Log(conf.C().Harbor.Scheme)
	t.Log(conf.C().Harbor.Username)
	t.Log(conf.C().Harbor.Password)
	t.Log(conf.C().Harbor.Host)
	t.Log(conf.C().Harbor.CaCerts)
	t.Log(conf.C().Harbor.GetHarborCaCerts())
	t.Log(conf.C().Harbor.InitHarborClient())
	t.Log(conf.C().Prometheus)
	t.Log(conf.C().Prometheus.Addr())
	t.Log(conf.C().Prometheus.GetPromAPI())
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
	t.Log(conf.C().Harbor.Scheme)
	t.Log(conf.C().Harbor.Username)
	t.Log(conf.C().Harbor.Password)
	t.Log(conf.C().Harbor.Host)
	t.Log(conf.C().Harbor.CaCerts)
	t.Log(conf.C().Harbor.GetHarborCaCerts())
	t.Log(conf.C().Harbor.InitHarborClient())
	t.Log(conf.C().Prometheus)
	t.Log(conf.C().Prometheus.Addr())
	t.Log(conf.C().Prometheus.GetPromAPI())
}
