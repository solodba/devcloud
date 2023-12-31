package impl

import (
	"github.com/solodba/devcloud/mkube/apps/pod"
	"github.com/solodba/devcloud/mkube/apps/pvc"
	"github.com/solodba/devcloud/mkube/apps/statefulset"
	"github.com/solodba/devcloud/mkube/conf"
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
	statefulset.UnimplementedRPCServer
	col       *mongo.Collection
	clientSet *kubernetes.Clientset
	podSvc    pod.Service
	pvcSvc    pvc.Service
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return statefulset.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("statefulsets")
	clientSet, err := conf.C().K8s.GetK8sConn()
	if err != nil {
		return err
	}
	i.clientSet = clientSet
	i.podSvc = apps.GetInternalApp(pod.AppName).(pod.Service)
	// 次奥, 麻痹的忘记初始化pvcSvc, 所以后面一直调用不到
	i.pvcSvc = apps.GetInternalApp(pvc.AppName).(pvc.Service)
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	statefulset.RegisterRPCServer(s, svc)
}

// 注册到Ioc中心
func init() {
	apps.RegistryInternalApp(svc)
	apps.RegistryGrpcApp(svc)
}
