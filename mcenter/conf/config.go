package conf

import (
	"fmt"
	"sync"

	"github.com/solodba/devcloud/tree/main/mcenter/common/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	c *Config
)

func C() *Config {
	if c == nil {
		logger.L().Panic().Msgf("please initial global config")
	}
	return c
}

type Http struct {
	Host string `toml:"host" env:"HTTP_HOST"`
	Port int    `toml:"port" env:"HTTP_PORT"`
}

type Grpc struct {
	Host string `toml:"host" env:"GRPC_HOST"`
	Port int    `toml:"port" env:"GRPC_PORT"`
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Http *Http  `toml:"http"`
	Grpc *Grpc  `toml:"grpc"`
}

type MongoDB struct {
	Username   string   `toml:"username" env:"MONGODB_USERNAME"`
	Password   string   `toml:"password" env:"MONGODB_PASSWORD"`
	Endpoints  []string `toml:"endpoints" env:"MONGODB_ENDPOINTS" envSeparator:","`
	Database   string   `toml:"database" env:"MONGODB_DATABASE"`
	AuthSource string   `toml:"authsource" env:"MONGODB_AUTHSOURCE"`
	lock       sync.Mutex
	client     *mongo.Client
}

type Config struct {
	App     *App     `toml:"app"`
	MongoDB *MongoDB `toml:"mongodb"`
}

func NewDefaultHttp() *Http {
	return &Http{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func NewDefaultGrpc() *Grpc {
	return &Grpc{
		Host: "127.0.0.1",
		Port: 8081,
	}
}

func NewDefaultApp() *App {
	return &App{
		Name: "mcenter",
		Http: NewDefaultHttp(),
		Grpc: NewDefaultGrpc(),
	}
}

func NewDefaultMongoDB() *MongoDB {
	return &MongoDB{
		Endpoints:  []string{"127.0.0.1:27017"},
		Database:   "mcenter",
		AuthSource: "mcenter",
	}
}

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
