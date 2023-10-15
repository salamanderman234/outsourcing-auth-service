package domain

import "errors"

var (
	// entity error
	ErrDuplicateKey = errors.New("duplicate entries")
	ErrMissingRequiredField = errors.New("missing required field(s)")
	ErrInvalidCreds = errors.New("invalid credentials")
	ErrRecordNotFound = errors.New("no record was found by given query")
	ErrTokenNotValid = errors.New("token is not valid")
	ErrTokenIsExpired = errors.New("token is expires")
	ErrConversionDataType = errors.New("cannot convert data type")
	ErrCreateToken = errors.New("cannot create authentication token")
	ErrPolicies = errors.New("sufficient access to this action")
	// response type
	ResponseSuccess = "success"
	ResponseDuplicateEntries = "duplicate entries error"
	ResponseValidationErr = "validation error"
	ResponseUnauthorizeErr = "unauthorize user error"
	ResponseForbiddenErr = "invalid access error"
	ResponseNotFoundErr = "not found error"
	ResponseBadRequest = "request error"
	ResponseTokenErr = "token error"
	ResponseServerErr = "internal server error"
)