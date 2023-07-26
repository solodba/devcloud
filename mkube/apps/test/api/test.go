package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/mcube/logger"
)

// 测试
func (h *handler) QueryTest(r *restful.Request, w *restful.Response) {
	attrToken := r.Attribute(token.ATTRIBUTE_TOKEN_KEY)
	if attrToken != "" {
		tk := attrToken.(*token.Token)
		logger.L().Info().Msgf("令牌: %s", tk)
	}
	w.WriteEntity("test")
}
