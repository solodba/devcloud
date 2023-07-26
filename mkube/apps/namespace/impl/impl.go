package impl

import (
	"github.com/solodba/devcloud/mkube/apps/namespace"
	"github.com/solodba/devcloud/mkube/conf"
	"github.com/solodba/mcube/apps"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	namespace.UnimplementedRPCServer
	clientSet *kubernetes.Clientset
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return namespace.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	clientSet, err := conf.C().K8s.GetK8sConn()
	if err != nil {
		return err
	}
	i.clientSet = clientSet
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	namespace.RegisterRPCServer(s, svc)
}

// 注册到Ioc中心
func init() {
	apps.RegistryInternalApp(svc)
	apps.RegistryGrpcApp(svc)
}
