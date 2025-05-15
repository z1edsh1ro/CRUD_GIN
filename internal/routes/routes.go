package router

import (
	"main/internal/adapter/http"
	"main/internal/adapter/repository"
	"main/internal/application/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(mongoClient *mongo.Client) *gin.Engine {
	r := gin.Default()

	todoRepository := repository.NewTodoRepository(mongoClient)
	todoService := service.TodoService{Port: todoRepository}
	todoHandler := http.NewTodoHandler(todoService)

	api := r.Group("/api/todo")
	{
		api.GET("/", todoHandler.List)
		api.GET("/:id", todoHandler.Get)
		api.POST("/", todoHandler.Create)
		api.PUT("/:id", todoHandler.Update)
		api.DELETE("/:id", todoHandler.Delete)
	}

	return r
}
