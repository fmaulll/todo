package main

import (
	"os"
	"time"

	"github.com/fmaulll/todolist/controllers"
	"github.com/fmaulll/todolist/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://todo-production-dd3d.up.railway.app/"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/api/todos", controllers.Index)
	router.POST("/api/todos", controllers.Create)
	router.PATCH("/api/todos/:id", controllers.Update)
	router.DELETE("/api/todos", controllers.Delete)

	router.Run(":" + os.Getenv("PORT"))
}
