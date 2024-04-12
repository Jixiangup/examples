package server

import (
	"jixiangup.me/examples/gin/internal/api"
	"jixiangup.me/examples/gin/internal/config"
	"jixiangup.me/examples/gin/internal/datasource"
	"jixiangup.me/examples/gin/pkg/logger"
)

func main() {
	config.SetupBootstrap()
	logger.SetupLogger()
	datasource.SetupMySQL()
	api.SetupWebServer()
}

func Run() {
	main()
}
