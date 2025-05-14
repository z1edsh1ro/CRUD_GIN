package http

import (
	"fmt"
	"main/internal/application/service"
	"main/internal/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{Service: service}
}

func (handler *TodoHandler) List(context *gin.Context) {
	todos, err := handler.Service.GetAllTodo()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   todos,
	})
}

func (handler *TodoHandler) Get(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "ERROR ID IS EMPTY",
		})
		return
	}

	todo, err := handler.Service.GetTodo(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   todo,
	})
}

func (handler *TodoHandler) Create(context *gin.Context) {
	var newTodo model.Todo

	err := context.ShouldBindJSON(&newTodo)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	err = handler.Service.CreateTodo(newTodo)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  "CREATED",
		"message": "CREATE TODO SUCCESS",
	})
}

func (handler *TodoHandler) Update(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "ERROR",
			"message": "ERROR ID IS EMPTY",
		})
		return
	}

	var todoUpdate model.Todo

	err := context.ShouldBindJSON(&todoUpdate)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	_, err = handler.Service.UpdateTodo(id, todoUpdate)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "UPDATED",
		"message": fmt.Sprintf("UPDATE TODO ID: %s SUCCESS", id),
	})
}

func (handler *TodoHandler) Delete(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "ERROR ID IS EMPTY",
		})
		return
	}

	err := handler.Service.DeleteTodo(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "DELETED",
		"message": fmt.Sprintf("DELETE TODO ID: %s SUCCESS", id),
	})
}
