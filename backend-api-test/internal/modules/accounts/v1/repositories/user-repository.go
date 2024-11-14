package repositories

import (
	"backend-api-test/config"
	"backend-api-test/internal/modules/accounts/v1/models"
	"backend-api-test/pkg/logger"
	"errors"
	"github.com/google/uuid"
	"time"
)

type IUserRepository interface {
	GetByEmail(email string) (*models.User, error)
	GetByUUID(uuid string) (*models.User, error)
}

type userRepository struct {
	conf   *config.Config
	logger logger.Logger
}

var timeNow = time.Now()
var firstUserUUID, _ = uuid.Parse("01930b4e-516f-76c4-8334-b2f76be901b8")
var secondUserUUID, _ = uuid.Parse("01930b57-d031-7bdb-9969-57e5f34c1b0b")

var UserList []models.User = []models.User{
	models.User{
		ID:        1,
		UUID:      firstUserUUID,
		Email:     "fulan@gmail.com",
		FullName:  "fulan",
		Gender:    models.Male,
		Hash:      "8e293d7f527fca0c084a99b191e2b18a", // fulan123 -> md5HASH
		Photo:     "https://plus.unsplash.com/premium_photo-1666777247416-ee7a95235559?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Salt:      "",
		LastLogin: nil,
		CreatedAt: &timeNow,
		UpdatedAt: nil,
		DeletedAt: nil,
		IsDeleted: false,
	},
	models.User{
		ID:        2,
		UUID:      secondUserUUID,
		Email:     "fulana@gmail.com",
		FullName:  "fulan A",
		Gender:    models.Male,
		Hash:      "6d37c9df8e3edf9947e55f8da0da223f", // fulanb123 -> md5HASH
		Photo:     "https://plus.unsplash.com/premium_photo-1694819488591-a43907d1c5cc?q=80&w=1014&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Salt:      "",
		LastLogin: nil,
		CreatedAt: &timeNow,
		UpdatedAt: nil,
		DeletedAt: nil,
		IsDeleted: false,
	},
}

func getUserByUUID(uuid string) (*models.User, error) {
	for _, user := range UserList {
		if uuid == user.UUID.String() {
			return &user, nil
		}
	}
	return nil, errors.New("the credentials was wrong")
}

func getUserByEmail(email string) (*models.User, error) {
	for _, user := range UserList {
		if email == user.Email {
			return &user, nil
		}
	}
	return nil, errors.New("the credentials was wrong")
}

func NewUserRepository(conf *config.Config, logger logger.Logger) IUserRepository {
	return &userRepository{
		conf:   conf,
		logger: logger,
	}
}

func (u *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	getDataUser, errGetDataUser := getUserByEmail(email)
	if errGetDataUser != nil {
		return nil, errGetDataUser
	}
	user = *getDataUser
	return &user, nil
}

func (u *userRepository) GetByUUID(uuid string) (*models.User, error) {
	var user models.User
	getDataUser, errGetDataUser := getUserByUUID(uuid)
	if errGetDataUser != nil {
		return nil, errGetDataUser
	}
	user = *getDataUser
	return &user, nil
}
