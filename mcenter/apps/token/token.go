package token

import "time"

// Token初始化函数
func NewToken() *Token {
	now := time.Now()
	return &Token{
		IssueAt:          now.Unix(),
		AccessExpiredAt:  60,
		RefreshExpiredAt: 60 * 4,
		Status:           &Status{},
		Meta:             map[string]string{},
	}
}
