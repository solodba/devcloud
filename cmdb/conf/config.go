package conf

import (
	"fmt"
	"sync"

	"github.com/solodba/devcloud/mcenter/client/rpc"
	"github.com/solodba/mcube/logger"
	"gorm.io/gorm"
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

// 定义MySQL连接信息结构体
type MySQL struct {
	Username    string `toml:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" env:"MYSQL_PASSWORD"`
	Host        string `toml:"host" env:"MYSQL_HOST"`
	Port        int64  `toml:"port" env:"MYSQL_PORT"`
	DB          string `toml:"db" env:"MYSQL_DB"`
	MaxOpenConn int64  `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int64  `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int64  `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int64  `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`
	lock        sync.Mutex
	db          *gorm.DB
}

// 全局配置结构体
type Config struct {
	App     *App                         `toml:"app"`
	MySQL   *MySQL                       `toml:"MySQL"`
	Mcenter *rpc.McenterGrpcClientConfig `toml:"mcenter"`
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

// MySQL结构体构造函数
func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Username:    "root",
		Password:    "123456",
		Host:        "127.0.0.1",
		Port:        3306,
		DB:          "test",
		MaxOpenConn: 50,
		MaxIdleConn: 10,
	}
}

// Config初始化函数
func NewDefaultConfig() *Config {
	return &Config{
		App:     NewDefaultApp(),
		MySQL:   NewDefaultMySQL(),
		Mcenter: rpc.NewDefaultMcenterGrpcClientConfig(),
	}
}

func (h *Http) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (g *Grpc) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}
