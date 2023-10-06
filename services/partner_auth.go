package service

import (
	"context"
	"strconv"
	"time"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
	"golang.org/x/crypto/bcrypt"
)

type partnerAuthService struct {
	repo domain.Repository
	role string
}

func NewAuthService(repo domain.Repository) domain.PartnerAuthService {
	return &partnerAuthService {
		repo: repo,
		role: "partner",
	}
}

func(a *partnerAuthService) Login(ctx context.Context, creds domain.Entity) (domain.AuthTokens, error) {
	var credsModel model.Partner
	var authTokens domain.AuthTokens
	err := helper.ConvertEntityToModel(creds, &credsModel)
	if err != nil {
		return authTokens, err
	}
	if credsModel.Password == nil {
		return authTokens, domain.ErrInvalidCreds
	}
	password := *credsModel.Password
	credsModel.Password = nil
	data, err := a.repo.Get(ctx, credsModel.Search)
	if err != nil {
		return authTokens, err
	}
	result := data.([]model.Partner)
	if len(result) != 1 {
		return authTokens, domain.ErrInvalidCreds
	}

	user := result[0]
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		return authTokens, domain.ErrInvalidCreds
	}
	token, err := helper.CreateToken(user.Email, user.Name, &user.Avatar, &a.role, user.ID, domain.TokenExpiresAt)
	if err != nil {
		return authTokens, err
	}
	refresh, err := helper.CreateToken(nil, nil, nil, &a.role, user.ID, domain.TokenRefreshExpiresAt)
	if err != nil {
		return authTokens, err
	}
	authTokens.Refresh = refresh
	authTokens.Token = token
	return authTokens, nil
}

func(a *partnerAuthService) Register(ctx context.Context, data domain.Entity) (domain.AuthTokens, error) {
	var dataModel model.Partner
	var authTokens domain.AuthTokens

	err := helper.ConvertEntityToModel(data, &dataModel)
	if err != nil {
		return authTokens, err
	}
	new, err := a.repo.Create(ctx, &dataModel)
	if err != nil {
		return authTokens, err
	}
	user := new.(*model.Partner)
	token, err := helper.CreateToken(user.Email, user.Name, &user.Avatar, &a.role, user.ID, domain.TokenExpiresAt)
	if err != nil {
		return authTokens, err
	}
	refresh, err := helper.CreateToken(nil, nil, nil, &a.role, user.ID, domain.TokenRefreshExpiresAt)
	if err != nil {
		return authTokens, err
	}
	authTokens.Refresh = refresh
	authTokens.Token = token
	return authTokens, nil
}

func(a *partnerAuthService) CheckTokenValid(token string) (domain.JWTClaims, error) {
	claims, err := helper.VerifyToken(token)
	if err != nil {
		return domain.JWTClaims{}, err
	}

	return claims, nil
}

func(a *partnerAuthService) RenewToken(ctx context.Context, refreshToken string) (domain.AuthTokens, error) {
	var authTokens domain.AuthTokens
	claims, err := a.CheckTokenValid(refreshToken)
	if err != nil {
		return authTokens, err
	}
	id,_ := strconv.Atoi(claims.ID)
	userData, err := a.repo.FindById(ctx, uint(id), &model.Partner{})
	if err != nil {
		return authTokens, err
	}
	user := userData.(*model.Partner)
	newToken, err := helper.CreateToken(user.Email, user.Name, &user.Avatar, &a.role, user.ID, domain.TokenExpiresAt)
	if err != nil {
		return authTokens, err
	}
	refresh, err := helper.CreateToken(nil, nil, nil, &a.role, user.ID, time.Now().Add(claims.ExpiresAt.Time.Sub(time.Now())))
	if err != nil {
		return authTokens, err
	}
	authTokens.Refresh = refresh
	authTokens.Token = newToken
	return authTokens, nil
}