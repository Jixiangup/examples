package api

import (
	"github.com/gin-gonic/gin"
	"jixiangup.me/examples/gin/internal/middleware"
)

// 注册账户相关路由
func registerAccounts(engine *gin.Engine) {
	registerAccountWhitelist(engine)
	registerAuthAccounts(engine)
}

// 注册账户白名单路由
func registerAccountWhitelist(engine *gin.Engine) {
	engine.GET("/account/:id", accountController.AccountDetail)
}

// 注册需要鉴权的账户路由
func registerAuthAccounts(engine *gin.Engine) {
	// 如果请求 `GET /use/account/:id`, `DELETE /use/account/:id` 则需要鉴权
	engine.Use(middleware.BearerAuthorizationMiddleware())
	{
		engine.GET("/use/account/:id", accountController.AccountDetail)
		engine.DELETE("/use/account/:id", accountController.AccountDetail)
	}

	// 或者 如果使用统一前缀鉴权
	// 如果请求 `GET /auth/account/:id, DELETE /auth/account/:id` 则需要鉴权
	authGroup := engine.Group("/account")
	authGroup.Use(middleware.BearerAuthorizationMiddleware())
	authGroup.DELETE("/:id", accountController.AccountDetail)
	authGroup.POST("", accountController.EditorAccount)
	authGroup.PUT("", accountController.EditorAccount)
}
