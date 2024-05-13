package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryID    uint
	Category      Category
	ImageProducts []ImageProduct
	Title         string
	Slug          string
	Sku           string
	Quantity      int
	OriginalPrice uint
	SalePrice     uint
	Status        bool
	Description   string
}
