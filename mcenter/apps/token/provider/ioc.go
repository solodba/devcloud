package provider

import (
	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
	"github.com/solodba/mcube/logger"
)

// 实例映射
var (
	issuerAppStore = map[token.GRANT_TYPE]Issuer{}
)

// 注册实例类
func Registry(k token.GRANT_TYPE, v Issuer) {
	if _, ok := issuerAppStore[k]; ok {
		logger.L().Panic().Msgf("issuer %v has registried ioc center, please don't registry again", k)
	}
	issuerAppStore[k] = v
	logger.L().Info().Msgf("issuer %v registry ioc center success", k)
}

// 获取实例类
func Get(k token.GRANT_TYPE) any {
	if v, ok := issuerAppStore[k]; ok {
		return v
	}
	logger.L().Panic().Msgf("issuer %v don't registry ioc center, please registry first", k)
	return nil
}

// 初始化实例类映射
func Init() error {
	for _, v := range issuerAppStore {
		if err := v.Conf(); err != nil {
			return err
		}
		logger.L().Info().Msgf("issuer initial object config success")
	}
	return nil
}
