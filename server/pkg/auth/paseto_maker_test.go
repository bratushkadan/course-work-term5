package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/chacha20poly1305"
)

const (
	tokenSecPaseto = "5bdb9986-2cbe-4e5e-a8b6-293e3ca9"
)

func TestPasetoMakerNot32CharLenSec(t *testing.T) {
	// arrange
	const myCoolSecret = "dark brandon"

	// act
	maker, err := NewPasetoMaker(myCoolSecret)

	// assert
	// key of len() != 32 is not possible for Paseto implementation of Maker
	if !errors.Is(err, ErrInvalidKeySize) {
		t.Error(err)
	}
	if maker != nil {
		t.Errorf("maker shouldn't be created when provided secret length is not of length %d", chacha20poly1305.KeySize)
	}
}

func TestPasetoMaker(t *testing.T) {
	// arrange
	var maker Maker
	maker, _ = NewPasetoMaker(tokenSecPaseto)

	username := "vladislav"
	duration := time.Minute

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	// act
	var token string
	token, _ = maker.CreateToken(username, 0, duration)
	userClaims, err := maker.VerifyToken(token)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, userClaims)
	require.NotZero(t, userClaims.ID)
	require.Equal(t, username, userClaims.Username)
	claimsIssuedAt, _ := userClaims.GetIssuedAt()
	claimsExpiresAt, _ := userClaims.GetExpirationTime()
	require.WithinDuration(t, issuedAt, claimsIssuedAt.Time, time.Second)
	require.WithinDuration(t, expiresAt, claimsExpiresAt.Time, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	// arrange
	var maker Maker
	maker, _ = NewJWTMaker(tokenSecPaseto)

	// act
	var token string
	token, _ = maker.CreateToken("bratushkadan", 0, -time.Minute)
	userClaim, err := maker.VerifyToken(token)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, ErrTokenExpired.Error())
	require.Nil(t, userClaim)
}
