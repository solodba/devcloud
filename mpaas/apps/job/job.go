package job

import "github.com/solodba/mcube/pb/meta"

// Job初始化函数
func NewJob(req *CreateJobRequest) *Job {
	return &Job{
		Meta: meta.NewMeta(),
		Job:  req,
	}
}

// Job默认初始化函数
func NewDefaultJob() *Job {
	return &Job{
		Meta: meta.NewMeta(),
		Job:  NewCreateJobRequest(),
	}
}

// JobSet初始化函数
func NewJobSet() *JobSet {
	return &JobSet{
		Items: make([]*JobSetItem, 0),
	}
}

// JobSet添加方法
func (j *JobSet) AddItems(items ...*JobSetItem) {
	j.Items = append(j.Items, items...)
}
