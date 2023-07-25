package cronjob

import (
	"github.com/solodba/mcube/pb/meta"
)

// CronJob初始化结构体
func NewCronJob(req *CreateCronJobRequest) *CronJob {
	return &CronJob{
		Meta:    meta.NewMeta(),
		CronJob: req,
	}
}
