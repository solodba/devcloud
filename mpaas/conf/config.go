package conf

import (
	"fmt"
	"sync"

	"github.com/solodba/mcube/logger"
	"go.mongodb.org/mongo-driver/mongo"
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

// 全局配置结构体
type Config struct {
	App     *App     `toml:"app"`
	MongoDB *MongoDB `toml:"mongodb"`
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
		Name: "mcenter",
		Http: NewDefaultHttp(),
		Grpc: NewDefaultGrpc(),
	}
}

// MongoDB初始化函数
func NewDefaultMongoDB() *MongoDB {
	return &MongoDB{
		Endpoints:  []string{"127.0.0.1:27017"},
		Database:   "mpaas",
		AuthSource: "mpaas",
	}
}

// Config初始化函数
func NewDefaultConfig() *Config {
	return &Config{
		App:     NewDefaultApp(),
		MongoDB: NewDefaultMongoDB(),
	}
}

func (h *Http) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (g *Grpc) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}
