package middleware

import (
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, valid := ParseJWT(c)
		if !valid {
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("unique_id", claims.UniqueID)
		c.Next()
	}
}
