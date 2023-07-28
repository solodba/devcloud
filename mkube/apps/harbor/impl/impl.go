package impl

import (
	"net/http"

	"github.com/solodba/devcloud/mkube/apps/harbor"
	"github.com/solodba/devcloud/mkube/conf"
	"github.com/solodba/mcube/apps"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	harbor.UnimplementedRPCServer
	col    *mongo.Collection
	harbor *conf.Harbor
	client *http.Client
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return harbor.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("harbors")
	i.harbor = conf.C().Harbor
	i.client, err = conf.C().Harbor.InitHarborClient()
	if err != nil {
		return err
	}
	return nil
}

// 实现Ioc中心RegistryHandler方法
func (i *impl) RegistryHandler(s *grpc.Server) {
	harbor.RegisterRPCServer(s, svc)
}

// 注册到Ioc中心
func init() {
	apps.RegistryInternalApp(svc)
	apps.RegistryGrpcApp(svc)
}
