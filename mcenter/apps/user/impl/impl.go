package impl

import (
	"github.com/solodba/devcloud/tree/main/mcenter/apps"
	"github.com/solodba/devcloud/tree/main/mcenter/apps/user"
	"github.com/solodba/devcloud/tree/main/mcenter/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	user.UnimplementedRPCServer
	col *mongo.Collection
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return user.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("users")
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	user.RegisterRPCServer(s, svc)
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
	apps.RegistryGrpcApp(svc)
}
