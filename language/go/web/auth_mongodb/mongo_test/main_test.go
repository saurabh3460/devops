package main

import (
	"auth-mongo/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSignupHandler(t *testing.T) {
	First_name := "saurabh"
	Last_name := "sy"
	Password := "aberf56"
	Email := "joker@dc.com"
	Phone := "(512) 522-1302"
	User_type := "USER"

	data := models.User{
		First_name: &First_name,
		Last_name:  &Last_name,
		Password:   &Password,
		Email:      &Email,
		Phone:      &Phone,
		User_type:  &User_type,
	}
	router := SetUpRouter()
	w := httptest.NewRecorder()
	jsonValue, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/users/signup", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

}
