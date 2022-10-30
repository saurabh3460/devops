package routes

import (
	controller "auth-mongo/controllers"

	"github.com/gin-gonic/gin"
)

var handler = map[string][]gin.HandlerFunc{
	"signup": {controller.Signup},
	"Login":  {controller.Login},
}

func AuthRoutes(routes *gin.Engine) {
	routes.POST("user/signup", handler["signup"]...)
	routes.POST("user/login", handler["login"]...)
}
