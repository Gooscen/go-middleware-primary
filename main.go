package main

import (
	"go-middleware-primary/config"
	"go-middleware-primary/router"
)

func main() {
	config.InitDB() // 初始化数据库
	config.DB.FirstOrCreate(&config.User{
		Username: "admin",
		Password: "123456",
	})
	r := router.SetupRouter() // 初始化路由
	r.Run(":8080")
}
