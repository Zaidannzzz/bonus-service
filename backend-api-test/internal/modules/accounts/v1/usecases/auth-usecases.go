package usecases

import (
	"backend-api-test/config"
	"backend-api-test/internal/modules/accounts/v1/models"
	"backend-api-test/pkg/jwt"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/http"
)

type IAuthUseCase interface {
	Login(ctx context.Context, data models.AuthLogin) (*models.Cookies, int32, error)
}

type authUseCase struct {
	userUseCase IUserUseCase
	config      config.Config
}

func NewAuthUseCase(userUseCase IUserUseCase, config config.Config) IAuthUseCase {
	return &authUseCase{
		userUseCase: userUseCase,
		config:      config,
	}
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (useCase *authUseCase) Login(ctx context.Context, data models.AuthLogin) (*models.Cookies, int32, error) {
	var cookies models.Cookies

	dataUser, _, errDataUser := useCase.userUseCase.GetByEmail(*data.Email)
	if errDataUser != nil {
		return nil, http.StatusInternalServerError, errDataUser
	}

	passwordHash := GetMD5Hash(*data.Password)
	if dataUser.Hash != passwordHash {
		return nil, http.StatusInternalServerError, errors.New("the credentials was wrong")
	}

	userUUIDString := dataUser.UUID.String()
	accessToken, errCreateAccessToken := jwt.GenerateToken(
		dataUser.Email,
		"at",
		userUUIDString,
		&useCase.config,
	)
	if errCreateAccessToken != nil {
		return nil, int32(http.StatusInternalServerError), errCreateAccessToken
	}

	cookies.AccessToken = accessToken

	return &cookies, http.StatusOK, nil
}
