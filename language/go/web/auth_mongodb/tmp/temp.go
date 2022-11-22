package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.RegisteredClaims
}

var secret []byte

func main() {
	claims := SignedDetails{
		Email:      "saurabh@gmail.com",
		First_name: "saurabh",
		Last_name:  "sy",
		Uid:        "88",
		User_type:  "editor",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))),
		},
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)

	// parse, validate
	Ptoken, _ := jwt.ParseWithClaims(token, &SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
	// fmt.Println(Ptoken, err)
	claims, ok := Ptoken.Claims.(SignedDetails)
	/*
		Ptoken.Claims.(SignedDetails) because claims can be of type Claims interface and it can be anytyhing which impliments Claim interface so through assertion we can check and get values

	*/
	fmt.Println(Ptoken, ok)
}
