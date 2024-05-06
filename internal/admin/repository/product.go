package repository

import (
	"context"
	"errors"
	"fmt"
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
		CategoryID:    req.CategoryID,
		Title:         title,
		Slug:          strings.ReplaceAll(title, " ", "-"),
		Sku:           strconv.Itoa(rand.IntN(9999)),
		Quantity:      int(req.Quantity),
		OriginalPrice: req.OriginalPrice,
		SalePrice:     req.SalePrice,
		Status:        status,
		Description:   req.Description,
	}

	p.Db.Begin()

	result := p.Db.Create(&product)
	if result.RowsAffected <= 0 {
		p.Db.Rollback()
		return entity.Product{}, errors.New("failed while creating new product")
	}
	//store images in db
	for _, image := range images {
		imgResult := p.Db.Table("image_product").Create(&entity.ImageProduct{
			ProductID: product.ID,
			Path:      image,
		}).Error
		if imgResult != nil {
			p.Db.Rollback()
			return entity.Product{}, errors.New("failed while creating new product")
		}
	}
	p.Db.Commit()

	return product, nil
}

func (p Product) GetAll(ctx context.Context) ([]entity.Product, error) {

	var products []entity.Product
	//todo: load category data
	result := p.Db.Preload("Category").Where("status", true).Find(&products)
	fmt.Println("\n\t --- result ", result.RowsAffected, "\n\t --- products ", products)
	if result.RowsAffected <= 0 {
		fmt.Println("\n\t --- not found ")
		return products, errors.New("no records found")
	}
	return products, nil
}
