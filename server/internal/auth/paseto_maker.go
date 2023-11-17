package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var (
	ErrInvalidKeySize = errors.New("invalid key size")
)

func (m *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	userClaim, err := NewUserClaim(username, duration)
	if err != nil {
		return "", err
	}
	return m.paseto.Encrypt(m.symmetricKey, userClaim, nil)
}
func (m *PasetoMaker) VerifyToken(token string) (*UserClaim, error) {
	userClaim := &UserClaim{}

	err := m.paseto.Decrypt(token, m.symmetricKey, userClaim, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = userClaim.Valid()
	if err != nil {
		return nil, err
	}

	return userClaim, err
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("%w: must be exactly of length %d", ErrInvalidKeySize, chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}
