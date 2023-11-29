package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/google/uuid"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Id       int32  `json:"id"`
}

var (
	ErrTokenExpired = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

func (p *UserClaim) Valid() error {
	expirationTime, err := p.GetExpirationTime()
	if err != nil {
		return err
	}
	if time.Now().After(expirationTime.Time) {
		return fmt.Errorf("%w at %s", ErrTokenExpired, p.ExpiresAt.Format("2006-01-02 15:04:05"))
	}
	return nil
}

func NewUserClaim(username string, id int32, duration time.Duration) (*UserClaim, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	userClaim := &UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId.String(),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Username: username,
		Id:       id,
	}
	return userClaim, nil
}
