package service

import (
	"main/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	todo := model.Todo{
		Id:      0,
		Context: "TEST",
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": todo,
	})
}

func Get(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": id,
	})
}

func Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "CREATE",
	})
}

func Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "UPDATE",
	})
}

func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "DELETE",
	})
}
