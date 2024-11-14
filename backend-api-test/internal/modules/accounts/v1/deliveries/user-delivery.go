package deliveries

import (
	"backend-api-test/config"
	"backend-api-test/internal/modules/accounts/v1/usecases"
	api_response "backend-api-test/pkg/api-response"
	"backend-api-test/pkg/helpers"
	"backend-api-test/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type IUserDelivery interface {
	GetByUUID(ctx *gin.Context)
	GetUserProfile(ctx *gin.Context)
}

type userDelivery struct {
	userUseCase usecases.IUserUseCase
	conf        config.Config
	logger      logger.Logger
}

func NewUserDelivery(
	userUseCase usecases.IUserUseCase,
	conf config.Config,
	logger logger.Logger,
) IUserDelivery {
	return &userDelivery{
		userUseCase: userUseCase,
		conf:        conf,
		logger:      logger,
	}
}

func (delivery *userDelivery) GetByUUID(ctx *gin.Context) {
	uuidUser := ctx.Param("user_uuid")
	uuidUserValidation, UserValidation := uuid.Parse(uuidUser)
	if UserValidation != nil {
		errorMessages := errors.New("ups, invalid uuid pattern")
		delivery.logger.Error(errorMessages)
		ctx.SecureJSON(http.StatusBadRequest, api_response.BadRequest(errorMessages.Error()))
		return
	}

	userUUIDString := uuidUserValidation.String()
	dataUser, httpStatusGetByUUID, errGetByUUID := delivery.userUseCase.GetByUUID(userUUIDString)
	if errGetByUUID != nil {
		delivery.logger.Error(errGetByUUID)
		helpers.ResponseError(httpStatusGetByUUID, errGetByUUID, ctx)
		return
	}

	ctx.SecureJSON(http.StatusOK, api_response.Success(nil, dataUser))
	return
}

func (delivery *userDelivery) GetUserProfile(ctx *gin.Context) {
	userUUIDAccessToken := helpers.GetUserUUIDFromMiddleware(ctx)

	dataGetByUUID, httpStatusGetByUUID, errGetByUUID := delivery.userUseCase.GetByUUID(userUUIDAccessToken)
	if errGetByUUID != nil {
		delivery.logger.Error(errGetByUUID)
		helpers.ResponseError(httpStatusGetByUUID, errGetByUUID, ctx)
		return
	}

	ctx.SecureJSON(http.StatusOK, api_response.Success(nil, dataGetByUUID))
	return
}
