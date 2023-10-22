package domain

import "errors"

var (
	// entity error
	ErrDuplicateKey = errors.New("data already exists")
	ErrInvalidCreds = errors.New("invalid credentials")
	ErrRecordNotFound = errors.New("resource not found")
	ErrTokenNotValid = errors.New("token is not valid")
	ErrTokenIsExpired = errors.New("token is expires")
	ErrConversionDataType = errors.New("cannot convert data type")
	ErrCreateToken = errors.New("cannot create authentication token")
	ErrPolicies = errors.New("user dont have access to this resource")
	ErrBindAndValidation = errors.New("request does not complify the rules")
	ErrUserSessionNotFound = errors.New("login is required for this action")
	ErrRefreshCookieNotFound = errors.New("refresh token is required")
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