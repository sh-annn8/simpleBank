package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// 不同类型的错误
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invvalid")
)

// Payload 包含 token 中的 payload data
type Payload struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	IssueAt time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// 根据 username 和 时间 创建一个新的 token payload
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID: tokenID,
		Username: username,
		IssueAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
