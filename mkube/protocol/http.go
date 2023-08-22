package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/conf"
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
	// 接入到mcenter认证中间件, 提供grpc认证凭证
	// r.Filter(auth.NewHttpAuther(conf.C().Mcenter).AuthFunc)
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

// Http服务启动方法
func (s *HttpService) Start() error {
	apps.InitRestfulApps(s.PathPrefix(), s.r)
	// 通过grpc获取service id
	// serviceId, err := conf.C().GetServiceIdByClientId()
	// if err != nil {
	// 	return err
	// }
	// 动态注册路由信息到mcenter
	// es := endpoint.NewEndpointSet()
	// ws := s.r.RegisteredWebServices()
	// for i := range ws {
	// 	routes := ws[i].Routes()
	// 	for j := range routes {
	// 		ep := endpoint.NewDefaultEndpoint()
	// 		ep.Spec.Method = routes[j].Method
	// 		ep.Spec.Path = routes[j].Path
	// 		ep.Spec.Operation = routes[j].Operation
	// 		ep.Spec.Perm = true
	// 		ep.Spec.ServiceId = serviceId
	// 		isAuth := routes[j].Metadata["auth"]
	// 		if isAuth != nil {
	// 			ep.Spec.Auth = isAuth.(bool)
	// 		}
	// 		es.AddItems(ep)
	// 	}
	// }
	// 调用grpc把路由信息注册到mcenter中
	// err = conf.C().WriteEndpointSetToMcenter(es)
	// if err != nil {
	// 	return err
	// }
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: swagger.DocsMkube,
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
