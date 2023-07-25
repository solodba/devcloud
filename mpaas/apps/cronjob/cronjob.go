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

// CronJob默认初始化结构体
func NewDefaultCronJob() *CronJob {
	return &CronJob{
		Meta:    meta.NewMeta(),
		CronJob: NewCreateCronJobRequest(),
	}
}
