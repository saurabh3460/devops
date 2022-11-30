package controllers

import (
	"auth-mongo/database"
	"auth-mongo/helpers"
	"auth-mongo/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	validate            = validator.New()
	BadRequest          = http.StatusBadRequest
	InternalServerError = http.StatusInternalServerError
)
var userColl = database.OpenCollection(database.Client, "user")

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
		msg = "email or password is incorrect"
		check = false
	}
	return check, msg
}

func Signup(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(BadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(BadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	count, err := userColl.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		c.JSON(InternalServerError, gin.H{"error": "error occured while checking for the email"})
		log.Panic(err)
	}

	if count > 0 {
		c.JSON(InternalServerError, gin.H{"error": "this email already exists"})
	}

	count, err = userColl.CountDocuments(ctx, bson.M{"phone": user.Phone})
	// defer cancel()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this phone number already exists"})

	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := helpers.GenerateAllToken(*user.Email, *user.First_name, *user.Last_name, *user.User_type, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	resultInsertionNumber, insertErr := userColl.InsertOne(ctx, user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User item was not created"})
		return
	}
	c.JSON(http.StatusOK, resultInsertionNumber)
}

func Login(c *gin.Context) {

}

func GetUser(c *gin.Context) {

	// uid := c.Param("user_id")
	// if err := helpers.MatchUserTypeToUid(c, uid); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
}

func GetUsers(c *gin.Context) {

}
