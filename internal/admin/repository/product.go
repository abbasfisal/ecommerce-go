package repository

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"gorm.io/gorm"
	"math/rand/v2"
	"strconv"
	"strings"
)

type Product struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) Product {
	return Product{Db: db}
}

func (p Product) Store(ctx context.Context, req requests.CreateProductRequest, images []string) (entity.Product, error) {

	title := strings.TrimSpace(req.Title)
	status := false
	if req.Status == "on" {
		status = true
	}
	product := entity.Product{
		Title:         title,
		Slug:          strings.ReplaceAll(title, " ", "-"),
		Sku:           strconv.Itoa(rand.IntN(9999)),
		Quantity:      int(req.Quantity),
		OriginalPrice: req.OriginalPrice,
		SalePrice:     req.SalePrice,
		Status:        status,
		Description:   req.Description,
	}

	result := p.Db.Create(&product)
	if result.RowsAffected <= 0 {
		return entity.Product{}, result.Error
	}
	return product, nil
}
