package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
)

func CreateClaims(email *string, username *string ,avatar *string, group *string, id uint, expire time.Time) domain.JWTClaims {
	claims := domain.JWTClaims {
		RegisteredClaims: jwt.RegisteredClaims{
			ID: strconv.Itoa(int(id)),
			ExpiresAt: jwt.NewNumericDate(expire),
		},
		JWTPayload: domain.JWTPayload{
			Email: email,
			Username: username,
			Group: group,
			Avatar: avatar,
		},
	}
	return claims
}

func CreateToken(email *string, username *string ,avatar *string, group *string, id uint, expire time.Time) (string, error) {
	claims := CreateClaims(email, username, avatar, group, id, expire)
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
