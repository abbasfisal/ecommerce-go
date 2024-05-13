package repository

import (
	"context"
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"gorm.io/gorm"
)

type PublicRepository struct {
	Db *gorm.DB
}

func NewPublicRepository(db *gorm.DB) PublicRepository {
	return PublicRepository{
		Db: db,
	}
}

func (p PublicRepository) GetAllProducts(ctx context.Context, offset int, perPage int) (products []entity.Product, totalCount int64, error error) {

	p.Db.Model(&entity.Product{}).Count(&totalCount)

	if err := p.Db.Offset(offset).Limit(perPage).Find(&products).Error; err != nil {
		fmt.Println("\n-- error in repo", err)
		return products, totalCount, err
	}
	return products, totalCount, nil
}
