package tools

import (
	"github.com/solodba/devcloud/tree/main/mcenter/apps"
	_ "github.com/solodba/devcloud/tree/main/mcenter/apps/all"
	"github.com/solodba/devcloud/tree/main/mcenter/common/logger"
	"github.com/solodba/devcloud/tree/main/mcenter/conf"
)

func DevelopmentSet() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		logger.L().Panic().Msgf("load global config error, err: %s", err.Error())
	}
	err = apps.InitInternalApps()
	if err != nil {
		logger.L().Panic().Msgf("initial object config error, err: %s", err.Error())
	}
}
