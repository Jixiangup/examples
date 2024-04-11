package api

import (
	"github.com/gin-gonic/gin"
	"jixiangup.me/examples/gin/internal/controller"
	"jixiangup.me/examples/gin/internal/datasource"
	"jixiangup.me/examples/gin/internal/repository"
)

var (
	accountController *controller.AccountController
)

func init() {

	accountController = controller.NewAccountController(repository.NewAccountRepository(datasource.MySQLInstance))
}

func SetupRoutes(engine *gin.Engine) {
	engine.GET("/account/:id", accountController.GetUserById)
}
