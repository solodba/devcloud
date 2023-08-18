package controller_test

import (
	"testing"
	"time"
)

func TestAuditLogSaveController(t *testing.T) {
	go c.Run(ctx)
	time.Sleep(300 * time.Second)
	err := c.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
