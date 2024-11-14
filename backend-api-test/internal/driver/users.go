package driver

import (
	"backend-api-test/internal/modules/accounts/v1/deliveries"
	"backend-api-test/internal/modules/accounts/v1/repositories"
	"backend-api-test/internal/modules/accounts/v1/usecases"
)

var (
	UserRepository = repositories.NewUserRepository(Conf, AppLogger)
	UserUseCase    = usecases.NewUserUseCase(UserRepository)
	UserDelivery   = deliveries.NewUserDelivery(UserUseCase, *Conf, AppLogger)
)
