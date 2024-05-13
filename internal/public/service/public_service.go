package service

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"github.com/abbasfisal/ecommerce-go/internal/public/contract"
)

const PerPage = 20

type PublicService struct {
	repo contract.PublicRepo
}

func NewPublicService(repo contract.PublicRepo) PublicService {
	return PublicService{
		repo: repo,
	}
}

func (p PublicService) GetProducts(ctx context.Context, page int) (products []entity.Product, totalCount int64, error error) {
	//todo: get from env variable

	offset := (page - 1) * PerPage
	return p.repo.GetAllProducts(ctx, offset, PerPage)
}
