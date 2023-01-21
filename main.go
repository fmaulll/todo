package main

import (
	"os"

	"github.com/fmaulll/todolist/controllers"
	"github.com/fmaulll/todolist/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://todo-production-dd3d.up.railway.app"}

	router.Use(cors.New(config))

	router.GET("/api/todos", controllers.Index)
	router.POST("/api/todos", controllers.Create)
	router.PATCH("/api/todos/:id", controllers.Update)
	router.DELETE("/api/todos", controllers.Delete)

	router.Run(":" + os.Getenv("PORT"))
}
