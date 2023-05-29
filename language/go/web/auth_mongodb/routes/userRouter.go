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
	routes.GET("/users", userHandlers["users"]...)
	routes.GET("/user/:user_id", userHandlers["user"]...)
}
