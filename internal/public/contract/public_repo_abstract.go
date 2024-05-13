package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type PublicRepo interface {
	GetAllProducts(ctx context.Context, offset, perPage int) (products []entity.Product, totalCount int64, error error)
}
