package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string
	PhoneNumber string
	Password    string
}
