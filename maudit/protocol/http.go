package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/maudit/conf"
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	"github.com/solodba/mcube/swagger"
)

// Http服务结构体
type HttpService struct {
	r   *restful.Container
	srv *http.Server
	c   *conf.Config
}

// Http服务结构体初始化
func NewHttpService() *HttpService {
	r := restful.DefaultContainer
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedDomains: []string{"*"},
		AllowedMethods: []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		CookiesAllowed: false,
		Container:      r,
	}
	r.Filter(cors.Filter)
	srv := &http.Server{
		Addr:              conf.C().App.Http.Addr(),
		Handler:           r,
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	return &HttpService{
		r:   r,
		srv: srv,
		c:   conf.C(),
	}
}

// 获取PathPrefix方法
func (h *HttpService) PathPrefix() string {
	return fmt.Sprintf("/%s/api", h.c.App.Name)
}

// Http服务启动方法 swagger文档对应每个模块
func (s *HttpService) Start() error {
	apps.InitRestfulApps(s.PathPrefix(), s.r)
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: swagger.DocsMaudit,
		DefinitionNameHandler: func(name string) string {
			if name == "state" || name == "sizeCache" || name == "unknownFields" {
				return ""
			}
			return name
		},
	}
	s.r.Add(restfulspec.NewOpenAPIService(config))
	logger.L().Info().Msgf("Get the API using http://%s%s", s.c.App.Http.Addr(), config.APIPath)
	logger.L().Info().Msgf("http service start success, listen address: %s", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logger.L().Info().Msgf("http service is stopped")
		}
		return fmt.Errorf("start http service error, %s", err.Error())
	}
	return nil
}

// Http服务停止方法
func (s *HttpService) Stop() error {
	logger.L().Info().Msgf("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(30))
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
