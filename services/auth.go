package service

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"path"
	"strconv"
	"time"

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
	// conversion
	if err := helper.ConvertEntityToModel(creds, credsModel); err != nil {
		return authTokens, domain.ErrConversionDataType
	}
	// set password to nil
	password := credsModel.GetPasswordField()
	credsModel.SetPasswordField(nil)	

	// get user
	data, err := a.repo.Get(ctx, credsModel.SearchQuery)
	if errors.Is(err, domain.ErrRecordNotFound) {
		return authTokens, domain.ErrInvalidCreds
	}
	if err != nil {
		return authTokens, err
	}
	result := data
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
	group := credsModel.GetGroupName()
	authTokens, err = helper.CreatePairTokenFromModel(user, group)
	if err != nil {
		return authTokens, domain.ErrCreateToken
	}
	return authTokens, nil
}

func(a *authService) Register(ctx context.Context, data domain.AuthEntity) (domain.AuthTokens, error) {
	var authTokens domain.AuthTokens
	dataModel := data.GetCorrespondingAuthModel()
	// data conversion
	err := helper.ConvertEntityToModel(data, dataModel)
	if err != nil {
		return authTokens, domain.ErrConversionDataType
	}
	// hashing password
	password := dataModel.GetPasswordField()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 1)
	hashedString := string(hashedPassword)
	dataModel.SetPasswordField(&hashedString)
	new, err := a.repo.Create(ctx, dataModel)
	if err != nil {
		return authTokens, err
	}
	// create token
	user := new.(domain.AuthModel)
	group := dataModel.GetGroupName()
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
	groupString := user.GetGroupName()
	authTokens, err = helper.CreatePairTokenFromModel(user, groupString)
	if err != nil {
		return authTokens, domain.ErrCreateToken
	}
	return authTokens, nil
}

func (a *authService) GenerateResetPasswordToken(ctx context.Context, email string, obj domain.AuthEntity) (error) {
	model := obj.GetCorrespondingAuthModel()
	err := helper.ConvertEntityToModel(obj, model)
	if err != nil {
		return domain.ErrConversionDataType
	}
	model.SetUsernameField(&email)
	result, err := a.repo.Get(ctx, model.SearchQuery)
	if err != nil {
		return nil
	}
	password := result[0].(domain.AuthModel).GetPasswordField()	
	group := model.GetGroupName()
	token, err  := helper.CreateResetPasswordToken(&group, &email, password, time.Now().Add(time.Duration(30) * time.Minute))
	if err != nil {
		return domain.ErrCreateToken
	}
	templatePath := path.Join("templates", "reset_password.html")
	tmplt, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	data := map[string]string{
		"token": token,
		"email": email,
		"group": group,
	}
	var tpl bytes.Buffer
	err = tmplt.Execute(&tpl, data)
	body := tpl.String()
	target := email
	subject := "Change password"
	go helper.SendMail(body, subject, target)
	return nil
}

func (a *authService) ResetPassword(ctx context.Context, token string, email string, newPassword string, obj domain.AuthEntity) error {
	model := obj.GetCorrespondingAuthModel()
	err := helper.ConvertEntityToModel(obj, model)
	if err != nil {
		return domain.ErrConversionDataType
	}
	model.SetUsernameField(&email)
	result, err := a.repo.Get(ctx, model.SearchQuery)
	if err != nil {
		return err
	}
	password := result[0].(domain.AuthModel).GetPasswordField()	
	_, err = helper.VerifyResetPasswordToken(token, password)
	if err != nil {
		return err
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(newPassword),1)
	hashedString := string(hashed)
	model.SetPasswordField(&hashedString)
	_, _, err = a.repo.Update(ctx, result[0].GetID(), model)
	if err != nil {
		return err
	}
	return nil
}