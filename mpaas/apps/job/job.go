package job

import "github.com/solodba/mcube/pb/meta"

// Job初始化函数
func NewJob(req *CreateJobRequest) *Job {
	return &Job{
		Meta: meta.NewMeta(),
		Job:  req,
	}
}
