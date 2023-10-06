package domain

import "errors"

var (
	ErrInvalidCreds = errors.New("invalid credentials")
	ErrRecordNotFound = errors.New("no record was found by given query")
	ErrTokenNotValid = errors.New("token is not valid")
	ErrTokenIsExpired = errors.New("token is expires")
)