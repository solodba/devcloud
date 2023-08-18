package broker

import (
	"context"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/solodba/devcloud/maudit/apps/audit"
)

const (
	TOPIC_NAME = "topic-audit"
)

// Producer客户端结构体
type Client struct {
	// kafka地址
	brokerAddress []string `toml:"address" env:"KAFKA_ADDRESS"`
	// kafka写接口
	writer *kafka.Writer
}

// Producer客户端构造函数
func NewClient(brokerAddress string) *Client {
	address := strings.Split(brokerAddress, ",")
	w := &kafka.Writer{
		Addr:                   kafka.TCP(address...),
		Topic:                  TOPIC_NAME,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: false,
	}
	return &Client{
		brokerAddress: address,
		writer:        w,
	}
}

// 客户端发送审计日志到kafka
func (c *Client) SendAuditLog(ctx context.Context, in *audit.AuditLog) error {
	msg := kafka.Message{
		Key:   []byte(in.ServiceId),
		Value: []byte(in.MustToJson()),
	}
	return c.writer.WriteMessages(ctx, msg)
}
