package router

import (
	handler "main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	todoHandler := handler.NewTodoHandler()

	api := r.Group("/api")
	api.GET("/", todoHandler.List)
	api.GET("/:id", todoHandler.Get)
	api.POST("/", todoHandler.Create)
	api.PUT("/", todoHandler.Update)
	api.DELETE("/:id", todoHandler.Delete)

	return r
}
