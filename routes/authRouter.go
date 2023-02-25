package routes

import (
	controller "github.com/brandonspitz/Go-JWT-Auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("user/login", controller.Login())
}
