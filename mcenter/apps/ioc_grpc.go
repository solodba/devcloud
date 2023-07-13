package apps

import (
	"github.com/solodba/devcloud/tree/main/mcenter/common/logger"
	"google.golang.org/grpc"
)

// 内部实例映射
var (
	grpcAppStore = map[string]GrpcIocObject{}
)

// 内部实例Ioc接口
type GrpcIocObject interface {
	IocObject
	RegistryHandler(*grpc.Server)
}

// Ioc内部实例注册函数
func RegistryGrpcApp(obj GrpcIocObject) {
	if _, ok := grpcAppStore[obj.Name()]; ok {
		logger.L().Panic().Msgf("grpc %s has registried ioc center, please don't registry again", obj.Name())
	}
	grpcAppStore[obj.Name()] = obj
	logger.L().Info().Msgf("grpc %s registry ioc center success", obj.Name())
}

// 从Ioc获取内部实例
func GetGrpcApp(objName string) any {
	if v, ok := grpcAppStore[objName]; ok {
		return v
	}
	logger.L().Panic().Msgf("grpc %s is not registry ioc center, please registry ioc center first", objName)
	return nil
}
