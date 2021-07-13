package model

import (

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(100);unique_index" json:"email"`
	FirstName    string `gorm:"size:100;not null" 			   json:"firstname"`
	LastName     string `gorm:"size:100;not null"              json:"lastname"`
	Password     string `gorm:"size:100;not null"              josn:"password"`
	ProfileImage string `gorm:"size:255"                       json:"profileimage"`
}
