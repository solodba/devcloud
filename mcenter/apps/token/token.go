package token

import (
	"fmt"
	"time"
)

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

// Token的校验
func (tk *Token) ValidateToken() error {
	expiredTime := tk.ExpiredTime()
	if time.Since(expiredTime) > 0 {
		return fmt.Errorf("令牌过期!")
	}
	return nil
}

// 获取过期时间
func (tk *Token) ExpiredTime() time.Time {
	return time.Unix(tk.IssueAt, 0).Add(time.Second * time.Duration(tk.AccessExpiredAt))
}
