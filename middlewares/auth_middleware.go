package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里只是一个简单的示例，实际项目中应该实现JWT验证等
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// 在实际项目中，这里应该验证token的有效性
		// 如果验证通过，可以将用户信息存储在上下文中
		// c.Set("user", userInfo)

		c.Next()
	}
}
