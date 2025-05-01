package router

import (
	"go-middleware-primary/controller"
	"go-middleware-primary/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RateLimitMiddleware())
	r.Use(middleware.HeaderMiddleware())

	// 登录接口（不需要鉴权）
	r.POST("/login", controller.Login)

	// 路由分组 + JWT鉴权
	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/profile", controller.Profile)
		api.GET("/data", controller.GetData)
	}

	return r
}
