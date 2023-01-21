package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fmaulll/todolist/models"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {

	var todos []models.Todo

	models.DB.Find(&todos)

	context.JSON(http.StatusOK, gin.H{"todos": todos})
}

func Create(context *gin.Context) {
	var todo models.Todo

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&todo)
	context.JSON(http.StatusCreated, gin.H{"todos": todo})
}

func Update(context *gin.Context) {
	var todo models.Todo

	id := context.Param("id")

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&todo).Where("id = ?", id).Updates(&todo).RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
}

func Delete(context *gin.Context) {
	var todo models.Todo

	var input struct {
		Id json.Number
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&todo, id).RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't delete todo!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
