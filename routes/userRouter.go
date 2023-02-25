package routes

import (
	controller "github.com/brandonspitz/Go-JWT-Auth/controllers"
	"github.com/brandonspitz/Go-JWT-Auth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
