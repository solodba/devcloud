package impl

import (
	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/devcloud/mkube/conf"
	"github.com/solodba/mcube/apps"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
)

var (
	svc = &impl{}
)

// 定义user的service的实例类
type impl struct {
	// 嵌套UnimplementedRPCServer结构体是因为它实现了RPCServer接口,最后相当于impl结构体实例实现了RPCServer接口
	metrics.UnimplementedRPCServer
	// 依赖MongoDB数据库
	col *mongo.Collection
	// k8s连接
	clientSet *kubernetes.Clientset
}

// 实现注册内部实例注册Name方法
func (i *impl) Name() string {
	return metrics.AppName
}

// 实现内部实例注册Conf方法
func (i *impl) Conf() error {
	// 获取MongoDB数据库连接
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	// 数据库操作metrics文档
	i.col = db.Collection("metrics")
	i.clientSet, err = conf.C().K8s.GetK8sConn()
	if err != nil {
		return err
	}
	return nil
}

// 实现grpc实例注册方法
func (i *impl) RegistryHandler(server *grpc.Server) {
	metrics.RegisterRPCServer(server, i)
}

// 注册metrics模块实例
func init() {
	// 注册grpc实例到IOC托管
	apps.RegistryGrpcApp(svc)
	// 注册内部实例到IOC托管
	apps.RegistryInternalApp(svc)
}
