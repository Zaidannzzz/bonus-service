package routes

import (
	"backend-api-test/internal/middleware"
	"backend-api-test/internal/modules/accounts/v1/deliveries"
	"github.com/gin-gonic/gin"
)

func InitAccountsRoutes(router *gin.Engine, userDelivery deliveries.IUserDelivery, authDelivery deliveries.IAuthDelivery, middleware middleware.IMiddleware) {
	accounts := router.Group("/accounts")
	v1Users := accounts.Group("/v1/users")
	v1Auth := accounts.Group("/v1/auth")

	v1Users.GET("/:user_uuid", middleware.AuthJwt(), userDelivery.GetByUUID)
	v1Users.GET("/profile", middleware.AuthJwt(), userDelivery.GetUserProfile)

	v1Auth.POST("/login", authDelivery.Login)
}
