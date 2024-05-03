package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type ProductService interface {
	Create(ctx context.Context, req requests.CreateProductRequest, images []string) (entity.Product, error)
}
