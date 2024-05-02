package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title  string
	Slug   string
	Image  string
	Status bool
}
