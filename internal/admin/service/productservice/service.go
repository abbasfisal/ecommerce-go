package productservice

import (
	"context"
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/admin/contract"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"os"
)

type ProductService struct {
	repo contract.ProductRepo
}

func NewService(repo contract.ProductRepo) ProductService {
	return ProductService{
		repo: repo,
	}
}

func (p ProductService) Create(ctx context.Context, req requests.CreateProductRequest, images []string) (entity.Product, error) {
	product, err := p.repo.Store(ctx, req, images)
	if err != nil {
		for _, imgPath := range images {
			if err := os.Remove(imgPath); err != nil {
				fmt.Println("\n\t --- product image removing failed", err.Error())
			}
			fmt.Println("\n\t -- product image successfully removed")
		}
		return entity.Product{}, err
	}
	return product, err
}
