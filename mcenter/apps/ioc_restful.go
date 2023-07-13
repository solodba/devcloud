package apps

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/tree/main/mcenter/common/logger"
)

// 内部实例映射
var (
	restfulAppStore = map[string]RestfulIocObject{}
)

// 内部实例Ioc接口
type RestfulIocObject interface {
	IocObject
	Version() string
	RegistryHandler(*restful.WebService)
}

// Ioc内部实例注册函数
func RegistryRestfulApp(obj RestfulIocObject) {
	if _, ok := restfulAppStore[obj.Name()]; ok {
		logger.L().Panic().Msgf("restful %s has registried ioc center, please don't registry again", obj.Name())
	}
	restfulAppStore[obj.Name()] = obj
	logger.L().Info().Msgf("restful %s registry ioc center success", obj.Name())
}

// 从Ioc获取内部实例
func GetRestfulApp(objName string) any {
	if v, ok := restfulAppStore[objName]; ok {
		return v
	}
	logger.L().Panic().Msgf("restful %s is not registry ioc center, please registry ioc center first", objName)
	return nil
}
