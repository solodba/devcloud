package impl

import (
	"github.com/solodba/devcloud/mcenter/apps/permission"
	"github.com/solodba/devcloud/mcenter/apps/policy"
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
	permission.UnimplementedRPCServer
	col *mongo.Collection
	// 引用policy服务
	svc policy.Service
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return permission.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("permissions")
	// 初始化引用的policy service
	i.svc = apps.GetInternalApp(policy.AppName).(policy.Service)
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	permission.RegisterRPCServer(s, svc)
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
	apps.RegistryGrpcApp(svc)
}
