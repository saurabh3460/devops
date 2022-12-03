package routes

import (
	controller "auth-mongo/controllers"

	"github.com/gin-gonic/gin"
)

var handler = map[string][]gin.HandlerFunc{
	"signup": {controller.Signup},
	"login":  {controller.Login},
}

func AuthRoutes(routes *gin.Engine) {
	routes.POST("/users/signup", handler["signup"]...)
	routes.POST("/users/login", handler["login"]...)
}
