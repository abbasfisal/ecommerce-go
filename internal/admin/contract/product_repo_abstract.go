package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type ProductRepo interface {
	Store(ctx context.Context, req requests.CreateProductRequest, images []string) (entity.Product, error)
	GetAll(ctx context.Context) ([]entity.Product, error)
}
