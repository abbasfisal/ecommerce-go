package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type CategoryRepo interface {
	Create(ctx context.Context, req *requests.CreateCategoryRequest, imageUploadedPath string) (entity.Category, error)
	List(ctx context.Context) ([]entity.Category, error)
	GetBy(ctx context.Context, ID string) (entity.Category, error)
}
