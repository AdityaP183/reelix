package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	controller "github.com/AdityaP183/reelix/server/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "alive", "message": "Backend is alive and running."})
	})

	router.POST("/movies", controller.AddMovie())
	router.GET("/movies", controller.GetMovies())
	router.GET("/movies/:imdb_id", controller.GetMovie())

	router.POST("/register", controller.RegisterUser())

	if err := router.Run(":8000"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
