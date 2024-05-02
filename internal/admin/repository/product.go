package repository

import (
	"gorm.io/gorm"
)

type Product struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) Product {
	return Product{Db: db}
}
