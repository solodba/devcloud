package meta

import (
	"time"

	"github.com/rs/xid"
)

// Meta初始化函数
func NewMeta() *Meta {
	return &Meta{
		Id:        xid.New().String(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
