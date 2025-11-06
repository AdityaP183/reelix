package routes

import (
	controller "github.com/AdityaP183/reelix/server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(router *gin.Engine) {
	router.GET("/movies", controller.GetMovies())

	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())

}
