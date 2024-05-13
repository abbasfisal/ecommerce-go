package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	PhoneNumber  string
	NationalCode string
	Password     string
	Type         string
	Address      Address
}
