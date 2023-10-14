package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
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
		domain.TokenExpiresAt,
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
		domain.TokenRefreshExpiresAt,
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
	signed, err := token.SignedString([]byte("asipa"))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func VerifyToken(token string)(domain.JWTClaims, error) {
	claims := domain.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return []byte("asipa"), nil
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		return claims, domain.ErrTokenIsExpired
	}
	if err != nil || !tkn.Valid {
		return claims, domain.ErrTokenNotValid
	}
	return claims, nil
}
