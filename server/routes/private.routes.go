package routes

import (
	controller "github.com/AdityaP183/reelix/server/controllers"
	"github.com/AdityaP183/reelix/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupPrivateRoutes(router *gin.Engine){
	router.Use(middleware.AuthMiddleware())

	router.POST("/movies", controller.AddMovie())
	router.GET("/movies/:imdb_id", controller.GetMovie())
}
