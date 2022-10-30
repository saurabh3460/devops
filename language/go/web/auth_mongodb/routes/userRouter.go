package routes

import (
	controller "auth-mongo/controllers"
	"auth-mongo/middleware"

	"github.com/gin-gonic/gin"
)

var userHandlers = map[string][]gin.HandlerFunc{
	"users": {controller.GetUsers},
	"user":  {controller.GetUser},
}

func UserRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate)
	routes.POST("/users", userHandlers["users"]...)
	routes.POST("/user/:user_id", userHandlers["user"]...)
}
