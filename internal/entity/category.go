package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Products []Product `gorm:"foreignKey:CategoryID"`
	Title    string
	Slug     string
	Image    string
	Status   bool
}
