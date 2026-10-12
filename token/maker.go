package token

import "time"

// token 接口，方便切换不同的服务
type Maker interface {
	// 为 username 创建新的 token
	CreateToken(username string, duration time.Duration) (string, error)

	// 验证 token 是否可用
	VerifyToken(token string) (*Payload, error)
}
