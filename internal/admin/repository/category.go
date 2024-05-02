package repository

import (
	"context"
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"gorm.io/gorm"
	"strings"
)

type Category struct {
	Db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) Category {
	return Category{Db: db}
}
func (c Category) Create(ctx context.Context, req *requests.CreateCategoryRequest, imageUploadedPath string) (entity.Category, error) {
	fmt.Println("\n ---- create repository -- ", req)
	var category entity.Category
	category.Title = strings.TrimSpace(req.Title)
	category.Slug = strings.ReplaceAll(category.Title, " ", "-")
	category.Image = imageUploadedPath

	if req.Status == "on" {
		category.Status = true
	} else {
		category.Status = false
	}
	result := c.Db.Create(&category)
	if result.RowsAffected <= 0 {
		return entity.Category{}, result.Error
	}

	return category, nil
}

func (c Category) List(ctx context.Context) ([]entity.Category, error) {
	var cats []entity.Category
	result := c.Db.Where("status", true).Find(&cats)
	if result.RowsAffected <= 0 {
		return cats, result.Error
	}
	return cats, nil
}
