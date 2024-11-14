package server

import (
	"backend-api-test/config"
	"backend-api-test/internal/driver"
	middleware2 "backend-api-test/internal/middleware"
	"backend-api-test/internal/modules/accounts/v1/routes"
	api_response "backend-api-test/pkg/api-response"
	"backend-api-test/pkg/helpers"
	"backend-api-test/pkg/helpers/databases/paginate"
	http_server "backend-api-test/pkg/http-server"
	"backend-api-test/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	logger logger.Logger
	conf   *config.Config
}

func InitServer(
	logger logger.Logger,
	conf *config.Config,
) *Server {
	return &Server{
		logger: logger,
		conf:   conf,
	}
}

func (s *Server) Run() {
	gin.SetMode(gin.ReleaseMode)
	// Paginate Start
	paginate.GetPaginationValueObject()

	router := gin.Default()
	router.RedirectTrailingSlash = false
	middleware := middleware2.InitMiddleware(s.conf, s.logger, driver.UserUseCase)

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.GlobalPanicHandler())

	// HTTP Server
	httpServer := http_server.Init(s.conf, router, http_server.Port(s.conf.Server.Port))

	routes.InitAccountsRoutes(router, driver.UserDelivery, driver.AuthDelivery, middleware)

	router.GET("/", func(ctx *gin.Context) {
		ctx.SecureJSON(http.StatusOK, api_response.StatusOK(nil, "the server was running"))
		return
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.SecureJSON(http.StatusNotFound, api_response.NotFound(nil, "Route"))
		return
	})

	router.NoMethod(func(ctx *gin.Context) {
		ctx.SecureJSON(http.StatusMethodNotAllowed, api_response.NotAllowed(nil, "Method"))
		return
	})

	/*
		Registration Custom Validation ENUM
	*/
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("Enum", helpers.Enum)
		_ = v.RegisterValidation("EpochTime", helpers.EpochTimeValidator)
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case server := <-interrupt:
		s.logger.Info("Server - Run - Signal: %s", server.String())
	case err := <-httpServer.Notify():
		s.logger.Errorf("Server - Run - httpServer.Notify: %s", err)
	}

	err := httpServer.Shutdown()
	if err != nil {
		s.logger.Errorf("Server - Run Error - httpServer.Shutdown: %s", err)
	}
}
