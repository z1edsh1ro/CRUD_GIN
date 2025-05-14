package http

import (
	"main/internal/domain/model"
	service "main/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		todoService: *service.NewTodoService(),
	}
}

func (h *TodoHandler) List(ctx *gin.Context) {
	todo := model.Todo{
		Id:      0,
		Content: "TEST",
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": todo,
	})
}

func (h *TodoHandler) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": id,
	})
}

func (h *TodoHandler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "CREATE",
	})
}

func (h *TodoHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "UPDATE",
	})
}

func (h *TodoHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	log.Info(id)

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "DELETE",
	})
}
