package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"ID"`
	Username  string    `json:"username"`
	ExpiredAt time.Time `json:"expired_at"`
	IssuedAt  time.Time `json:"issued_at"`
}

var ErrExpiredToken = errors.New("token has expired")

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
