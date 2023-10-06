package domain

import "errors"

var (
	// entity error
	ErrMissingRequiredField = errors.New("missing required field(s)")
	// 
	ErrInvalidCreds = errors.New("invalid credentials")
	ErrRecordNotFound = errors.New("no record was found by given query")
	ErrTokenNotValid = errors.New("token is not valid")
	ErrTokenIsExpired = errors.New("token is expires")
	// internal server error
	ErrConversionDataType = errors.New("cannot convert data type")
	ErrCreateToken = errors.New("cannot create authentication token")
)