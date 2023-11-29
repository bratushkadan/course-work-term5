package api

import (
	"errors"
)

var (
	ErrNotFound             = errors.New("not found")
	ErrBadCredentials       = errors.New("bad credentials")
	ErrFailedToGenAuthToken = errors.New("failed to generate auth token")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrForbidden            = errors.New("forbidden")
	ErrInternalServerError  = errors.New("internal server error")
)

type Impl struct {
}
