package conf

import (
	"fmt"
	"net/http"
	"sync"

	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/solodba/mcube/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/client-go/kubernetes"
)

// 全局配置参数
var (
	c *Config
)

func C() *Config {
	if c == nil {
		logger.L().Panic().Msgf("please initial global config")
	}
	return c
}

// Http结构体
type Http struct {
	Host string `toml:"host" env:"HTTP_HOST"`
	Port int    `toml:"port" env:"HTTP_PORT"`
}

// Grpc结构体
type Grpc struct {
	Host string `toml:"host" env:"GRPC_HOST"`
	Port int    `toml:"port" env:"GRPC_PORT"`
}

// App结构体
type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Http *Http  `toml:"http"`
	Grpc *Grpc  `toml:"grpc"`
}

// MongoDB结构体
type MongoDB struct {
	Username   string   `toml:"username" env:"MONGODB_USERNAME"`
	Password   string   `toml:"password" env:"MONGODB_PASSWORD"`
	Endpoints  []string `toml:"endpoints" env:"MONGODB_ENDPOINTS" envSeparator:","`
	Database   string   `toml:"database" env:"MONGODB_DATABASE"`
	AuthSource string   `toml:"authsource" env:"MONGODB_AUTHSOURCE"`
	lock       sync.Mutex
	client     *mongo.Client
}

// K8S配置结构体
type K8s struct {
	KubeConfig string `toml:"kubeconfig" env:"K8S_KUBECONFIG"`
	lock       sync.Mutex
	clientSet  *kubernetes.Clientset
}

// Plugin结构体
type Plugin struct {
	Provisionser string `toml:"provisionser" env:"PLUGIN_PROVISIONSER"`
}

// Harbor结构体
type Harbor struct {
	Scheme   string `toml:"scheme" env:"HARBOR_SCHEME"`
	Username string `toml:"username" env:"HARBOR_USERNAME"`
	Password string `toml:"password" env:"HARBOR_PASSWORD"`
	Host     string `toml:"host" env:"HARBOR_HOST"`
	CaCerts  string `toml:"cacerts" env:"HARBOR_CACERTS"`
	lock     sync.Mutex
	client   *http.Client
}

// Prometheus结构体
type Prometheus struct {
	Enable  bool   `toml:"enable" env:"PROMETHEUS_ENABLE"`
	Scheme  string `toml:"scheme" env:"PROMETHEUS_SCHEME"`
	Host    string `toml:"host" env:"PROMETHEUS_HOST"`
	Port    int32  `toml:"port" env:"PROMETHEUS_PORT"`
	lock    sync.Mutex
	promAPI promv1.API
}

// 全局配置结构体
type Config struct {
	App        *App        `toml:"app"`
	MongoDB    *MongoDB    `toml:"mongodb"`
	K8s        *K8s        `toml:"k8s"`
	Plugin     *Plugin     `toml:"plugin"`
	Harbor     *Harbor     `toml:"harbor"`
	Prometheus *Prometheus `toml:"prometheus"`
}

// Http初始化函数
func NewDefaultHttp() *Http {
	return &Http{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

// Grpc初始化函数
func NewDefaultGrpc() *Grpc {
	return &Grpc{
		Host: "127.0.0.1",
		Port: 8081,
	}
}

// App初始化函数
func NewDefaultApp() *App {
	return &App{
		Name: "mkube",
		Http: NewDefaultHttp(),
		Grpc: NewDefaultGrpc(),
	}
}

// MongoDB初始化函数
func NewDefaultMongoDB() *MongoDB {
	return &MongoDB{
		Endpoints:  []string{"127.0.0.1:27017"},
		Database:   "mkube",
		AuthSource: "mkube",
	}
}

// K8s初始化函数
func NewDefaultK8s() *K8s {
	return &K8s{
		KubeConfig: "etc/config",
	}
}

// Plugin初始化函数
func NewDefaultPlugin() *Plugin {
	return &Plugin{}
}

// Harbor初始化函数
func NewDefaultHarbor() *Harbor {
	return &Harbor{
		Scheme:   "https",
		Username: "test",
		Password: "test",
		Host:     "127.0.0.1",
	}
}

// Prometheus结构体初始化函数
func NewPrometheus() *Prometheus {
	return &Prometheus{
		Enable: false,
		Scheme: "http",
		Host:   "127.0.0.1",
		Port:   30090,
	}
}

// Config初始化函数
func NewDefaultConfig() *Config {
	return &Config{
		App:        NewDefaultApp(),
		MongoDB:    NewDefaultMongoDB(),
		K8s:        NewDefaultK8s(),
		Plugin:     NewDefaultPlugin(),
		Harbor:     NewDefaultHarbor(),
		Prometheus: NewPrometheus(),
	}
}

func (h *Http) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (g *Grpc) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}

func (p *Prometheus) Addr() string {
	return fmt.Sprintf("%s://%s:%d", p.Scheme, p.Host, p.Port)
}
