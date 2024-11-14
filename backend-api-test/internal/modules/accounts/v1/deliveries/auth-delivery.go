package deliveries

import (
	"backend-api-test/config"
	"backend-api-test/internal/modules/accounts/v1/models"
	"backend-api-test/internal/modules/accounts/v1/usecases"
	api_response "backend-api-test/pkg/api-response"
	"backend-api-test/pkg/helpers"
	"backend-api-test/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type IAuthDelivery interface {
	Login(ctx *gin.Context)
}

type authDelivery struct {
	authUseCase usecases.IAuthUseCase
	logger      logger.Logger
	conf        *config.Config
}

func InitAuthDelivery(
	authUseCase usecases.IAuthUseCase,
	logger logger.Logger,
	conf *config.Config,
) IAuthDelivery {
	return &authDelivery{
		authUseCase: authUseCase,
		logger:      logger,
		conf:        conf,
	}
}

func (delivery *authDelivery) Login(ctx *gin.Context) {
	var payload models.AuthLogin
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		delivery.logger.Error(err)
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorMessages := helpers.ErrorMessage(err)
			ctx.SecureJSON(http.StatusBadRequest, api_response.BadRequest(errorMessages[0]))
			return
		}
		ctx.SecureJSON(http.StatusBadRequest, api_response.BadRequest(err.Error()))
		return
	}

	validEmail := helpers.IsEmailValid(*payload.Email)
	if validEmail == false {
		errorMessage := errors.New("ups, e-mail doesn't valid")
		delivery.logger.Error(errorMessage)
		ctx.SecureJSON(http.StatusBadRequest, api_response.BadRequest(errorMessage.Error()))
		return
	}

	data, httpStatus, errorMessage := delivery.authUseCase.Login(ctx, payload)
	if errorMessage != nil {
		delivery.logger.Error(errorMessage)
		helpers.ResponseError(httpStatus, errorMessage, ctx)
		return
	}

	ctx.SecureJSON(http.StatusOK, api_response.Success(nil, data))
	return
}
