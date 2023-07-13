package apps

import "github.com/solodba/devcloud/tree/main/mcenter/common/logger"

// 内部实例映射
var (
	internalAppStore = map[string]IocObject{}
)

// 内部实例Ioc接口
type IocObject interface {
	Name() string
	Conf() error
}

// Ioc内部实例注册函数
func RegistryInternalApp(obj IocObject) {
	if _, ok := internalAppStore[obj.Name()]; ok {
		logger.L().Panic().Msgf("%s has registried ioc center, please don't registry again", obj.Name())
	}
	internalAppStore[obj.Name()] = obj
	logger.L().Info().Msgf("%s registry ioc center success", obj.Name())
}

// 从Ioc获取内部实例
func GetInternalApp(objName string) any {
	if v, ok := internalAppStore[objName]; ok {
		return v
	}
	logger.L().Panic().Msgf("%s is not registry ioc center, please registry ioc center first", objName)
	return nil
}
