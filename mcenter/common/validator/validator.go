package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/solodba/devcloud/tree/main/mcenter/common/logger"
)

var (
	validate *validator.Validate
)

func V() *validator.Validate {
	if validate == nil {
		logger.L().Panic().Msgf("please initial global validate")
	}
	return validate
}

func init() {
	validate = validator.New()
}
