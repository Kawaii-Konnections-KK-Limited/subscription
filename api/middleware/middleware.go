package middleware

import (
	"fmt"
	"net/http"

	"github.com/Kawaii-Konnections-KK-Limited/subscription/utils"
	"github.com/gin-gonic/gin"
)

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		fmt.Println(token)
		if token == "" {
			c.AbortWithStatus(401)
			return
		}
		if utils.VerifyToken(&token) {
			c.Set("token", &token)
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
		c.Abort()
	}
}
