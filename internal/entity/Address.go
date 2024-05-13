package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID      uint
	Floor       int
	Phase       string
	Number      string
	Block       string
	Description string
}
