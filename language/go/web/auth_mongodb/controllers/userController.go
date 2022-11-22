package controllers

import (
	"auth-mongo/database"
	"auth-mongo/helpers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	userColl *mongo.Collection = database.OpenCollection(database.Client, "user")
	validate                   = validator.New()
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check = false
	}
	return check, msg
}

func Signup(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func GetUser(c *gin.Context) {

	uid := c.Param("user_id")
	if err := helpers.MatchUserTypeToUid(c, uid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func GetUsers(c *gin.Context) {

}
