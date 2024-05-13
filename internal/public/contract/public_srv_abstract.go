package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type PublicSrv interface {
	GetProducts(ctx context.Context, perPage int) (products []entity.Product, totalCount int64, error error)
}
