package impl

import (
	"github.com/solodba/devcloud/mcenter/apps/endpoint"
	"github.com/solodba/devcloud/mcenter/conf"
	"github.com/solodba/mcube/apps"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	endpoint.UnimplementedRPCServer
	col *mongo.Collection
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return endpoint.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("endpoints")
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	endpoint.RegisterRPCServer(s, svc)
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
	apps.RegistryGrpcApp(svc)
}
