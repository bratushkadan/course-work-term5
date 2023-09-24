package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	tokenSec = "5bdb9986-2cbe-4e5e-a8b6-293e3ca9"
	username = "$cool_username"
)

func TestJWTMaker(t *testing.T) {
	// arrange
	duration := time.Minute
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	maker, _ := NewJWTMaker(tokenSec)
	token, _ := maker.CreateToken(username, duration)

	// act
	userClaims, err := maker.VerifyToken(token)
	claimsIssuedAt, _ := userClaims.GetIssuedAt()
	claimsExpiresAt, _ := userClaims.GetExpirationTime()

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, userClaims)
	require.NotZero(t, userClaims.ID)
	require.Equal(t, username, userClaims.Username)
	require.WithinDuration(t, issuedAt, claimsIssuedAt.Time, time.Second)
	require.WithinDuration(t, expiresAt, claimsExpiresAt.Time, time.Second)
}

func TestExpiredJWT(t *testing.T) {
	// arrange
	maker, _ := NewJWTMaker(tokenSec)
	token, _ := maker.CreateToken("bratushkadan", -1*time.Minute)

	// act
	userClaim, err := maker.VerifyToken(token)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, ErrTokenExpired.Error())
	require.Nil(t, userClaim)
}
