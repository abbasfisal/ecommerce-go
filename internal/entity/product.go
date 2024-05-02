package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryID    uint
	Title         string
	Slug          string
	Sku           string
	Quantity      int
	OriginalPrice uint
	SalePrice     uint
	Status        bool
	Description   string
}
