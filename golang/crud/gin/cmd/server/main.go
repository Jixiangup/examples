package server

import (
	"github.com/gin-gonic/gin"
	"jixiangup.me/examples/gin/internal/api"
	_ "jixiangup.me/examples/gin/internal/config"
	"jixiangup.me/examples/gin/internal/constants"
	_ "jixiangup.me/examples/gin/internal/datasource"
	"jixiangup.me/examples/gin/pkg/logger"
	"os"
)

func main() {

	gin.SetMode(os.Getenv(constants.GinMode))

	engine := gin.Default()
	api.SetupRoutes(engine)

	port := os.Getenv(constants.ServerPort)
	if port == "" {
		port = "6000"
	}

	if err := engine.Run(":" + port); err != nil {
		logger.Log.Error("Failed to start server", err)
	}
}

func Run() {
	main()
}
