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
	c = broker.NewClient("192.168.1.130:9092", "audit_log")
}
