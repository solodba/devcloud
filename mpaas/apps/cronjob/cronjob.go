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

// CronJobSet初始化函数
func NewCronJobSet() *CronJobSet {
	return &CronJobSet{
		Items: make([]*CronJobSetItem, 0),
	}
}

// CronJobSet添加方法
func (s *CronJobSet) AddItems(items ...*CronJobSetItem) {
	s.Items = append(s.Items, items...)
}
