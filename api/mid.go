package api

import (
	public "RoutePilot/api/public"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 校验 Cookie 或 Header Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		// 优先从 Header 中取
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// 如果 Header 没有，再从 Cookie 取
		if token == "" {
			cookie, err := c.Cookie("user")
			if err == nil {
				token = cookie
			}
		}

		// 如果两边都没有
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing token, please login",
			})
			return
		}

		// TODO: 在这里校验 token 是否有效（例如 JWT 验签、数据库验证等）
		if user, err := public.ParseToken(token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is invalid",
			})
			return
		} else {
			c.Set("user", user) // 举例：解析 token 得到用户信息
			c.Next()            // 继续执行后续处理
			return
		}

		// token 有效，可以把用户信息放进 context
		c.Next() // 继续执行后续处理

	}
}

// 伪 token 校验函数（实际项目中替换为 JWT 或数据库查询）
func validateToken(token string) bool {
	return token == "admin"
}
