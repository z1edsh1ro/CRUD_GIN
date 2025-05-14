package main

import (
	"main/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// setupRoutes(r)
	r.GET("/", service.List)
	r.GET("/:id", service.Get)
	r.POST("/", service.Create)
	r.PUT("/", service.Update)
	r.DELETE("/:id", service.Delete)

	r.Run()
}
