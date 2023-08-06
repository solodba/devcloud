package prometheus

// PromMetrics结构体初始化函数
func NewPromMetrics(metrics string) *PromMetrics {
	return &PromMetrics{
		Metrics: metrics,
	}
}
