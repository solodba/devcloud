package impl

import (
	"github.com/solodba/devcloud/mkube/apps/svc"
	"github.com/solodba/devcloud/mkube/conf"
	"github.com/solodba/mcube/apps"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
)

var (
	imp = &impl{}
)

// 业务实现类
type impl struct {
	svc.UnimplementedRPCServer
	col       *mongo.Collection
	clientSet *kubernetes.Clientset
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return svc.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("services")
	clientSet, err := conf.C().K8s.GetK8sConn()
	if err != nil {
		return err
	}
	i.clientSet = clientSet
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	svc.RegisterRPCServer(s, i)
}

// 注册到Ioc中心
func init() {
	apps.RegistryInternalApp(imp)
	apps.RegistryGrpcApp(imp)
}
