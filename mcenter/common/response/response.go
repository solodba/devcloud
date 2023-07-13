package response

// 返回错误数据结构
type Fail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 返回成功数据结构体
type Success struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// 错误结构体初始化函数
func NewFail(code int, msg string) *Fail {
	return &Fail{
		Code:    code,
		Message: msg,
	}
}

// 正确结构体初始化函数
func NewSuccess(code int, data any) *Success {
	return &Success{
		Code: code,
		Data: data,
	}
}
