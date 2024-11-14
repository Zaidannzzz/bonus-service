package middleware

import (
	"backend-api-test/config"
	"backend-api-test/internal/modules/accounts/v1/usecases"
	api_response "backend-api-test/pkg/api-response"
	"backend-api-test/pkg/helpers"
	"backend-api-test/pkg/jwt"
	"backend-api-test/pkg/logger"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type IMiddleware interface {
	CORSMiddleware() gin.HandlerFunc
	GlobalPanicHandler() gin.HandlerFunc
	AuthJwt() gin.HandlerFunc
}

type middleware struct {
	conf                *config.Config
	authorizationHeader string
	logger              logger.Logger
	useCaseUser         usecases.IUserUseCase
}

func InitMiddleware(
	conf *config.Config,
	logger logger.Logger,
	useCaseUser usecases.IUserUseCase,
) IMiddleware {
	return &middleware{
		conf:                conf,
		authorizationHeader: "Authorization",
		logger:              logger,
		useCaseUser:         useCaseUser,
	}
}

func (m *middleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Auth-User-Id, sec-ch-ua, sec-ch-ua-mobile, sec-ch-ua-platform, user-agent, additionalheaders, referer")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, HEAD, PATCH, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (m *middleware) GlobalPanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic
				log.Println("Panic occurred:", r)
				message := fmt.Sprintf("Panic occurred: %v", r)
				c.Set("panic", true)

				helpers.ResponseError(500, errors.New(message), c)
			}
		}()
		// Continue to the next middleware or handler
		c.Next()
	}
}

func (m *middleware) AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(m.authorizationHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)

		if len(token) == 0 || token == "" {
			errorMessage := fmt.Sprintf("Token is empty")
			c.SecureJSON(http.StatusNotImplemented, api_response.ServerNotImplemented(errorMessage))
			c.Abort()
			return
		}

		validateToken, err := jwt.ValidateToken(token, m.conf.JWT.SecretKey)

		if err != nil {
			errorMessage := fmt.Sprintf("%v", err)
			c.SecureJSON(http.StatusUnauthorized, api_response.UnAuthorized(nil, errorMessage))
			c.Abort()
			return
		}

		roleUUID, errExtract := jwt.ExtractTokenRoleID(validateToken)
		if errExtract != nil {
			errorMessage := fmt.Sprintf("%v", err)
			c.SecureJSON(http.StatusUnauthorized, api_response.UnAuthorized(nil, errorMessage))
			c.Abort()
			return
		}

		userUUID, errExtractUserUUID := jwt.ExtractTokenUserUUID(validateToken)
		if errExtractUserUUID != nil {
			errorMessage := fmt.Sprintf("%v", errExtractUserUUID)
			c.SecureJSON(http.StatusUnauthorized, api_response.UnAuthorized(nil, errorMessage))
			c.Abort()
			return
		}

		dataGetUserByUUID, httpStatusGetUserByUUID, errGetUserByUUID := m.useCaseUser.GetByUUID(*userUUID)
		if errGetUserByUUID != nil {
			helpers.ResponseError(int32(httpStatusGetUserByUUID), errGetUserByUUID, c)
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("role_uuid", *roleUUID)
		c.Set("user_uuid", dataGetUserByUUID.UUID.String())
		c.Set("user_id", dataGetUserByUUID.ID)
		c.Next()

	}
}
