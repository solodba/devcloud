package impl

import (
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/solodba/devcloud/mkube/apps/prometheus"
	"github.com/solodba/devcloud/mkube/conf"
	"github.com/solodba/mcube/apps"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svc = &impl{}
)

// 定义user的service的实例类
type impl struct {
	// 嵌套UnimplementedRPCServer结构体是因为它实现了RPCServer接口,最后相当于impl结构体实例实现了RPCServer接口
	prometheus.UnimplementedRPCServer
	// 依赖MongoDB数据库
	col *mongo.Collection
	// Prometheus客户端API
	promapi promv1.API
	// Prometheus结构体变量
	prom *conf.Prometheus
}

// 实现注册内部实例注册Name方法
func (i *impl) Name() string {
	return prometheus.AppName
}

// 实现内部实例注册Conf方法
func (i *impl) Conf() error {
	// 获取MongoDB数据库连接
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		return err
	}
	// 数据库操作prometheus文档
	i.col = db.Collection("prometheus")
	// 初始化Prometheus API客户端
	promapi, err := conf.C().Prometheus.GetPromAPI()
	if err != nil {
		return err
	}
	i.promapi = promapi
	// 初始化Prometheus结构体
	i.prom = conf.C().Prometheus
	return nil
}

// 实现grpc实例注册方法
func (i *impl) RegistryHandler(server *grpc.Server) {
	prometheus.RegisterRPCServer(server, i)
}

// 注册metrics模块实例
func init() {
	// 注册grpc实例到IOC托管
	apps.RegistryGrpcApp(svc)
	// 注册内部实例到IOC托管
	apps.RegistryInternalApp(svc)
}
