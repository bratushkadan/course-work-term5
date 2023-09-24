package auth

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	// arrange
	const pass = "f"

	// act
	hashedPass, err := HashPassword(pass)

	// assert
	if err != nil {
		t.Error(fmt.Errorf("error hashing password: %w", err))
	}
	if len(hashedPass) == 0 {
		t.Error(errors.New("error hashing password: length of the hashed password is 0"))
	}
}

func TestHashPasswordNotTooLong(t *testing.T) {
	// arrange
	pass := strings.Repeat("a", 73)

	// act
	hashedPass, err := HashPassword(pass)

	// assert
	if hashedPass != "" {
		t.Errorf("hashed password string must be empty: can't hash password of length >72")
	}
	if !errors.Is(err, bcrypt.ErrPasswordTooLong) {
		t.Error("error must be %w, but got %w", bcrypt.ErrPasswordTooLong, err)
	}
}

func TestMatchHashedPasswordMatches(t *testing.T) {
	// arrange
	const pass = "mysecretpassword"
	var hashedPass string

	// act
	hashedPass, _ = HashPassword(pass)
	matchingErr := MatchPassword(hashedPass, pass)

	// assert
	if matchingErr != nil {
		t.Error(matchingErr)
	}
}

func TestMatchHashedPasswordNotMatches(t *testing.T) {
	// arrange
	const pass = "mysecretpassword"
	const pass2 = "myothersecretpassword"
	var hashedPass string

	// act
	hashedPass, _ = HashPassword(pass)
	matchingErr := MatchPassword(hashedPass, pass2)

	// assert
	if !errors.Is(matchingErr, bcrypt.ErrMismatchedHashAndPassword) {
		t.Error(matchingErr)
	}
}
