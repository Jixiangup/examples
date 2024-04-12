package api

import (
	"github.com/gin-gonic/gin"
	"jixiangup.me/examples/gin/internal/middleware"
)

func registerAccounts(engine *gin.Engine) {
	registerAccountWhitelist(engine)
	registerAuthAccounts(engine)
}

func registerAccountWhitelist(engine *gin.Engine) {
	engine.GET("/account/:id", accountController.AccountDetail)
}

func registerAuthAccounts(engine *gin.Engine) {
	// 如果请求 `GET /use/account/:id`, `DELETE /use/account/:id` 则需要鉴权
	engine.Use(middleware.BearerAuthorizationMiddleware())
	{
		engine.GET("/use/account/:id", accountController.AccountDetail)
		engine.DELETE("/use/account/:id", accountController.AccountDetail)
	}

	// 或者 如果使用统一前缀鉴权
	// 如果请求 `GET /auth/account/:id, DELETE /auth/account/:id` 则需要鉴权
	authGroup := engine.Group("/auth")
	authGroup.Use(middleware.BearerAuthorizationMiddleware())
	authGroup.GET("/account/:id", accountController.AccountDetail)
	authGroup.DELETE("/account/:id", accountController.AccountDetail)
}
