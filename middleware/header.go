package middleware

import "github.com/gin-gonic/gin"

func HeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Powered-By", "Gin-Gonic")
		c.Next()
	}
}
