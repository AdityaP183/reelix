package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/AdityaP183/reelix/server/routes"
)

func main() {
	router := gin.Default()

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "alive", "message": "Backend is alive and running."})
	})
	routes.SetupPublicRoutes(router)
	routes.SetupPrivateRoutes(router)

	if err := router.Run(":8000"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
