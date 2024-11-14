package usecases

import (
	"backend-api-test/internal/modules/accounts/v1/models"
	"backend-api-test/internal/modules/accounts/v1/repositories"
	"net/http"
)

type IUserUseCase interface {
	GetByEmail(email string) (*models.User, int32, error)
	GetByUUID(uuid string) (*models.User, int32, error)
}

type userUseCase struct {
	userRepo repositories.IUserRepository
}

func NewUserUseCase(userRepo repositories.IUserRepository) IUserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (useCase *userUseCase) GetByEmail(email string) (*models.User, int32, error) {
	user, err := useCase.userRepo.GetByEmail(email)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	return user, http.StatusOK, nil
}

func (useCase *userUseCase) GetByUUID(uuid string) (*models.User, int32, error) {
	user, err := useCase.userRepo.GetByUUID(uuid)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	return user, http.StatusOK, nil
}
