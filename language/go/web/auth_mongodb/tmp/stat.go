package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {

	// expireDuration := time.Hour * -20 // Negative duration of 20 hours
	expireDuration := time.Minute * -1 // Negative duration of 20 hours

	expireTime := time.Now().Add(expireDuration)

	ExpiresAt := jwt.NewNumericDate(expireTime)

	fmt.Println(ExpiresAt.Before(time.Now().Local()))

	fmt.Println(expireTime, time.Now().Local(), expireDuration)

}
