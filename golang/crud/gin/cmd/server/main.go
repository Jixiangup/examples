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
	if err == nil {
		// 如果服务启动成功，那么就一直阻塞主线程
		select {}
	}
}
