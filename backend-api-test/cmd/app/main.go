package main

import (
	"backend-api-test/internal/driver"
	"backend-api-test/internal/server"
	"log"
)

func main() {
	conf := driver.Conf
	if driver.ErrConf != nil {
		log.Fatalf("Loading config: %v", driver.ErrConf)
	}
	appLogger := driver.AppLogger
	appLogger.InitLogger()
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s, Port:%v, SSL: %v",
		conf.Server.AppVersion,
		conf.Logger.Level,
		conf.Server.Mode,
		conf.Server.Port,
		conf.Server.SSL,
	)
	appLogger.Infof("Success parsed config: %#v", conf.Server.AppVersion)
	appLogger.Infof("Application running on Mode: %v", conf.Server.Mode)

	s := server.InitServer(appLogger, conf)
	s.Run()
}
