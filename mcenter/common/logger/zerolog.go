package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// 全局日志参数
var (
	logger *zerolog.Logger
)

func L() *zerolog.Logger {
	if logger == nil {
		log.Panic("please initial global logger")
	}
	return logger
}

// 初始化全局日志方法
func init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	l := zerolog.New(output).With().Timestamp().Caller().Logger()
	logger = &l
}
