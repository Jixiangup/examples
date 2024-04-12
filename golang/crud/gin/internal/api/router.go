package api

import (
	"github.com/gin-gonic/gin"
	"jixiangup.me/examples/gin/internal/constants"
	"jixiangup.me/examples/gin/internal/controller"
	"jixiangup.me/examples/gin/internal/datasource"
	"jixiangup.me/examples/gin/internal/middleware"
	"jixiangup.me/examples/gin/internal/repository"
	"jixiangup.me/examples/gin/internal/service"
	"jixiangup.me/examples/gin/pkg/logger"
	"os"
)

var (
	accountController *controller.AccountController
)

func setupControllers() {
	accountController = controller.NewAccountController(service.NewAccountService(repository.NewAccountRepository(datasource.MySQLInstance)))
}

func SetupRoutes() *gin.Engine {
	gin.SetMode(os.Getenv(constants.GinMode))
	engine := gin.Default()

	// 注册中间件
	engine = registerMiddleware(engine)

	// 注册控制器
	setupControllers()

	// 注册路由
	registerRouting(engine)

	port := os.Getenv(constants.ServerPort)
	if port == "" {
		port = "5266"
	}
	if err := engine.Run(":" + port); err != nil {
		logger.Log.Error("Failed to start server", err)
	}
	return engine
}

func registerMiddleware(engine *gin.Engine) *gin.Engine {
	// 注册异常处理中间件
	engine.Use(middleware.ErrorHandler())
	return engine
}

func registerRouting(engine *gin.Engine) {
	// 注册账户相关路由
	registerAccounts(engine)
}

func registerAccounts(engine *gin.Engine) {
	engine.GET("/account/:id", accountController.AccountDetail)
}
