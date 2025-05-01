package main

import (
	"go-middleware-primary/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
