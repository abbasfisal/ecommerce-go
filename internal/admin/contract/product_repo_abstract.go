package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type ProductRepo interface {
	Create(ctx context.Context) (entity.Product, error)
}
