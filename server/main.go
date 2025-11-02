package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/status", func(ctx *gin.Context) {
		ctx.String(200, "Backend is alive and running.")
	})

	if err := router.Run(":8000"); err != nil{
		fmt.Println("Failed to start server", err)
	}
}
