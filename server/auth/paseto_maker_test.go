package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	// arrange
	var maker Maker
	maker, _ = NewPasetoMaker(tokenSec)

	username := "bratushakdan"
	duration := time.Minute

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	// act
	var token string
	token, _ = maker.CreateToken(username, duration)
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
	maker, _ = NewJWTMaker(tokenSec)

	// act
	var token string
	token, _ = maker.CreateToken("bratushkadan", -time.Minute)
	userClaim, err := maker.VerifyToken(token)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, ErrTokenExpired.Error())
	require.Nil(t, userClaim)
}
