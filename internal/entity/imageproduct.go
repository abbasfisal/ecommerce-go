package entity

import "gorm.io/gorm"

type ImageProduct struct {
	gorm.Model
	ProductID uint
	Path      string
}
