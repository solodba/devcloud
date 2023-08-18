package broker_test

import (
	"context"

	"github.com/solodba/devcloud/maudit/client/broker"
)

var (
	c   *broker.Client
	ctx = context.Background()
)

func init() {
	client := broker.NewClient("192.168.1.130:9092", broker.DEFAULT_TOPIC_NAME)
	c = client
}
