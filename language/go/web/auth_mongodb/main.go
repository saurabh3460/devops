package main

import (
	"auth-mongo/routes"

	"github.com/gin-gonic/gin"
)

const (
	HOST        = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "main"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	port := "8080"

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "access granted for api-1"})
	})

	router.Run(":" + port)
}
