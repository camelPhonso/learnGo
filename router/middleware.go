package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Auth-Go-Key")

		if len(authHeader) == 0 || authHeader != "ducks" {
			c.Writer.Header().Set("WWW-Authenticate","Basic")
			c.JSON(http.StatusUnauthorized, gin.H{"Error":"Authentication Header invalid or missing."})	
			c.Abort()
		}

		c.Next()	
	}
}