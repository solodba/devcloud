package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/deployment"
	"github.com/solodba/devcloud/mpaas/apps/pod"
	"github.com/solodba/devcloud/mpaas/conf"
	"github.com/solodba/mcube/apps"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	deployment.UnimplementedRPCServer
	col       *mongo.Collection
	clientSet *kubernetes.Clientset
	podSvc    pod.Service
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return deployment.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("deployments")
	clientSet, err := conf.C().K8s.GetK8sConn()
	if err != nil {
		return err
	}
	i.clientSet = clientSet
	i.podSvc = apps.GetInternalApp(pod.AppName).(pod.Service)
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	deployment.RegisterRPCServer(s, svc)
}

// 注册到Ioc中心
func init() {
	apps.RegistryInternalApp(svc)
	apps.RegistryGrpcApp(svc)
}
