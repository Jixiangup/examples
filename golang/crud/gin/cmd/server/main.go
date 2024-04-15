package server

import (
	"jixiangup.me/examples/gin/internal/api"
	"jixiangup.me/examples/gin/internal/config"
	"jixiangup.me/examples/gin/internal/datasource"
	"jixiangup.me/examples/gin/pkg/logger"
)

func exec() error {
	config.SetupBootstrap()
	logger.SetupLogger()
	datasource.SetupMySQL()
	err := api.SetupWebServer()
	return err
}

func Run() {
	err := exec()
	if err != nil {
		logger.Log.Error("Failed to start server %s", err.Error())
	}
}
