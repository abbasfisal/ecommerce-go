package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type CategoryService interface {
	StoreCategory(ctx context.Context, req *requests.CreateCategoryRequest, imageUploadedPath string) (entity.Category, error)
	GetAll(ctx context.Context) ([]entity.Category, error)
}
