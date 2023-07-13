package apps_test

import (
	"testing"

	"github.com/solodba/devcloud/tree/main/mcenter/apps"
)

const (
	AppName = "test"
)

type Test struct {
	Age int
}

type Service interface {
	GetName() int
}

func (t *Test) Name() string {
	return AppName
}

func (t *Test) Conf() error {
	t.Age = 100
	return nil
}

func (t *Test) GetName() int {
	return t.Age
}

func TestIoc(t *testing.T) {
	apps.RegistryInternalApp(&Test{})
	err := apps.InitInternalApps()
	if err != nil {
		t.Fatal(err)
	}
	svc := apps.GetInternalApp(AppName).(Service)
	age := svc.GetName()
	t.Log(age)
}
