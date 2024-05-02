package categoryservice

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/contract"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type Service struct {
	repo contract.CategoryRepo
}

func NewService(repo contract.CategoryRepo) Service {
	return Service{repo: repo}
}
func (s Service) StoreCategory(ctx context.Context, req *requests.CreateCategoryRequest, imageUploadedPath string) (entity.Category, error) {
	//todo: u can also check for any errors to remove the uploaded image from disk in service or in handler
	return s.repo.Create(ctx, req, imageUploadedPath)
}
