package page

// 添加计算OffSet方法
func (p *PageRequest) ComputeOffSet() {
	p.Offset = (p.PageNumber - 1) * p.PageSize
}

// PageRequest初始化函数
func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageNumber: 1,
		PageSize:   5,
	}
}
