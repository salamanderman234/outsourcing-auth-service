package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-auth-profile-service/config"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
)

var (
	SECRET = config.GetAppSecret()
)

func CreateClaims(username *string, name *string ,avatar *string, group *string, id uint, expire time.Time) domain.JWTClaims {
	claims := domain.JWTClaims {
		RegisteredClaims: jwt.RegisteredClaims{
			ID: strconv.Itoa(int(id)),
			ExpiresAt: jwt.NewNumericDate(expire),
		},
		JWTPayload: domain.JWTPayload{
			Username: username,
			Name: name,
			Group: group,
			Avatar: avatar,
		},
	}
	return claims
}

func CreatePairTokenFromModel(data domain.AuthModel, group string) (domain.AuthTokens, error) {
	var pairs domain.AuthTokens
	usernameField := data.GetUsernameField()
	identityField := data.GetIdentityField()
	avatarField := data.GetAvatarField()
	idField := data.GetID()
	groupField := group

	access, err := CreateToken(
		&usernameField, 
		&identityField, 
		&avatarField, 
		&groupField, 
		idField, 
		time.Now().Add(time.Duration(30) * time.Minute),
	)
	if err != nil {
		return pairs, err
	}
	refresh, err := CreateToken(
		nil, 
		nil, 
		nil, 
		&groupField, 
		idField, 
		time.Now().Add(time.Duration(72) * time.Hour),
	)
	if err != nil {
		return pairs, err
	}
	pairs.Refresh = refresh
	pairs.Access = access
	return pairs, nil
}

func CreateToken(username *string, name*string ,avatar *string, group *string, id uint, expire time.Time) (string, error) {
	claims := CreateClaims(username, name, avatar, group, id, expire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func CreateResetPasswordToken(group *string, email *string, password string, expire time.Time) (string, error) {
	empty := ""
	claims := CreateClaims(email, &empty, &empty, group, 0, expire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(password))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func VerifyResetPasswordToken(token string, key string) (domain.JWTClaims, error) {
	claims := domain.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return []byte(key), nil
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		return claims, domain.ErrTokenIsExpired
	}
	if err != nil || !tkn.Valid {
		return claims, domain.ErrTokenNotValid
	}
	return claims, nil
}

func VerifyToken(token string)(domain.JWTClaims, error) {
	claims := domain.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return []byte(SECRET), nil
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		return claims, domain.ErrTokenIsExpired
	}
	if err != nil || !tkn.Valid {
		return claims, domain.ErrTokenNotValid
	}
	return claims, nil
}
