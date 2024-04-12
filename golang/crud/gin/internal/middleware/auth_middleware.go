package middleware

import (
	"github.com/gin-gonic/gin"
	"jixiangup.me/examples/gin/internal/controller"
	error2 "jixiangup.me/examples/gin/internal/error"
	"strings"
)

// BearerAuthorizationMiddleware 验证 Bearer Token
// 该中间件用于验证请求头中的Authorization字段是否为Bearer Token
// 如果不是Bearer Token则返回401 Unauthorized
// 如果是Bearer Token但是验证失败则返回401 Unauthorized
// 如果是Bearer Token且验证成功则放行
// 该中间件适用于需要验证Token的接口
// 用例: curl -H "Authorization: Bearer testing" http://127.0.0.1:5266/use/account/1
func BearerAuthorizationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			controller.Error(context, error2.NewUnauthorized("Unauthorized"))
			return
		}

		// 获取值结果应该是 `Bearer ${token}`
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			controller.Error(context, error2.NewUnauthorized("Unauthorized"))
			return
		}

		token := parts[1]

		if !validateToken(token) {
			controller.Error(context, error2.NewUnauthorized("Unauthorized"))
			return
		}

		// 放行
		context.Next()
	}
}

// 假设的Token验证函数
func validateToken(token string) bool {
	// 这里应添加逻辑来验证JWT token的有效性
	// 例如使用github.com/dgrijalva/jwt-go库来解析和验证token
	// token写如redis或者数据库，然后在这里查询都是可以的
	return token == "testing" // 举例，实际应用中请替换为真实的验证逻辑
}
