package impl

import (
	"github.com/solodba/devcloud/mcenter/apps/policy"
	"github.com/solodba/devcloud/mcenter/apps/role"
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
	policy.UnimplementedRPCServer
	col *mongo.Collection
	// 引入role服务
	svc role.Service
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return policy.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("policys")
	i.svc = apps.GetInternalApp(role.AppName).(role.Service)
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	policy.RegisterRPCServer(s, svc)
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
	apps.RegistryGrpcApp(svc)
}
