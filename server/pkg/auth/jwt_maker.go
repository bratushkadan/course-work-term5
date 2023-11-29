package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

func (m *JWTMaker) CreateToken(username string, id int32, duration time.Duration) (string, error) {
	userClaim, err := NewUserClaim(username, id, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	return jwtToken.SignedString([]byte(m.secretKey))
}
func (m *JWTMaker) VerifyToken(token string) (*UserClaim, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &UserClaim{}, keyFunc)
	if err != nil {

		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}

		// if ok && errors.Is(, ErrTokenExpired) {
		// 	return nil, ErrTokenExpired
		// }
		return nil, err
	}

	userClaim, ok := jwtToken.Claims.(*UserClaim)
	if !ok {
		return nil, ErrInvalidToken
	}
	return userClaim, nil
}

const minSecretKeyLen = 32

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeyLen {
		return nil, fmt.Errorf("invalid secretKey size: must be at least %d characters", minSecretKeyLen)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}
