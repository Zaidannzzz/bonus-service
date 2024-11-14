package driver

import (
	"backend-api-test/internal/modules/accounts/v1/deliveries"
	"backend-api-test/internal/modules/accounts/v1/usecases"
)

var (
	AuthUseCase  = usecases.NewAuthUseCase(UserUseCase, *Conf)
	AuthDelivery = deliveries.InitAuthDelivery(AuthUseCase, AppLogger, Conf)
)
