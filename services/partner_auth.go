package service

import (
	"context"
	"reflect"
	"strconv"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo domain.Repository
}

func NewAuthService(repo domain.Repository) domain.AuthService {
	return &authService {
		repo: repo,
	}
}

func(a *authService) Login(ctx context.Context, creds domain.AuthEntity) (domain.AuthTokens, error) {
	var authTokens domain.AuthTokens
	credsModel := creds.GetCorrespondingAuthModel()
	// check required field
	if !creds.CheckRequiredLoginField() {
		return authTokens, domain.ErrMissingRequiredField
	}
	// conversion
	if err := helper.ConvertEntityToModel(creds, credsModel); err != nil {
		return authTokens, domain.ErrConversionDataType
	}
	// set password to nil
	password := credsModel.GetPasswordField()
	credsModel.SetPasswordField(nil)	

	// get user
	data, err := a.repo.Get(ctx, credsModel.SearchQuery)
	if err != nil {
		return authTokens, err
	}
	result := data.([]domain.Model)
	if len(result) != 1 {
		return authTokens, domain.ErrInvalidCreds
	}
	// comparing password
	user := result[0].(domain.AuthModel)
	err = bcrypt.CompareHashAndPassword([]byte(user.GetPasswordField()), []byte(password))
	if err != nil {
		return authTokens, domain.ErrInvalidCreds
	}
	// creating token
	group := reflect.TypeOf(user).Elem().Name()
	authTokens, err = helper.CreatePairTokenFromModel(user, group)
	if err != nil {
		return authTokens, domain.ErrCreateToken
	}
	return authTokens, nil
}

func(a *authService) Register(ctx context.Context, data domain.AuthEntity) (domain.AuthTokens, error) {
	var authTokens domain.AuthTokens
	dataModel := data.GetCorrespondingAuthModel()
	// check required field
	if !data.CheckRequiredRegisterField() {
		return authTokens, domain.ErrMissingRequiredField
	}
	// data conversion
	err := helper.ConvertEntityToModel(data, dataModel)
	if err != nil {
		return authTokens, domain.ErrConversionDataType
	}
	// hashing password
	password := dataModel.GetPasswordField()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 3)
	hashedString := string(hashedPassword)
	dataModel.SetPasswordField(&hashedString)
	new, err := a.repo.Create(ctx, dataModel)
	if err != nil {
		return authTokens, err
	}
	// create token
	user := new.(domain.AuthModel)
	group := reflect.TypeOf(user).Elem().Name()
	authTokens, err = helper.CreatePairTokenFromModel(user, group)
	if err != nil {
		return authTokens, domain.ErrCreateToken
	}
	return authTokens, nil
}

func(a *authService) CheckTokenValid(token string) (domain.JWTClaims, error) {
	claims, err := helper.VerifyToken(token)
	if err != nil {
		return domain.JWTClaims{}, err
	}

	return claims, nil
}

func(a *authService) RenewToken(ctx context.Context, refreshToken string, group domain.AuthEntity) (domain.AuthTokens, error) {
	var authTokens domain.AuthTokens
	claims, err := a.CheckTokenValid(refreshToken)
	if err != nil {
		return authTokens, err
	}
	id,_ := strconv.Atoi(claims.ID)
	userData, err := a.repo.FindById(ctx, uint(id), group.GetCorrespondingModel())
	if err != nil {
		return authTokens, domain.ErrTokenNotValid
	}
	user := userData.(domain.AuthModel)
	groupString := reflect.TypeOf(user).Elem().Name()
	authTokens, err = helper.CreatePairTokenFromModel(user, groupString)
	if err != nil {
		return authTokens, domain.ErrCreateToken
	}
	return authTokens, nil
}