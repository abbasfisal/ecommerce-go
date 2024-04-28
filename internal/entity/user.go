package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	PhoneNumber string
	Password    string
	Type        string
}
