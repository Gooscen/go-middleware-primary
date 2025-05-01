package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

var limiter = time.Tick(100 * time.Millisecond) // 每100ms处理1个请求

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		<-limiter
		c.Next()
	}
}
