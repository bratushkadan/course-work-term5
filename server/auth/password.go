package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hashedPassBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(hashedPassBytes), err
}

func MatchPassword(hashedPass, providedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(providedPass))
}
