package driver

import (
	"backend-api-test/config"
	"backend-api-test/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

var (
	_             = godotenv.Load()
	configPath    = config.GetConfigPath(os.Getenv("config"))
	Conf, ErrConf = config.GetConfig(configPath)
	AppLogger     = logger.InitApiLogger(Conf)
)
