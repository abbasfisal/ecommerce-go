package service

import "github.com/abbasfisal/ecommerce-go/internal/public/contract"

type PublicService struct {
	repo contract.PublicRepo
}

func NewPublicService(repo contract.PublicRepo) PublicService {
	return PublicService{
		repo: repo,
	}
}

func (p PublicService) Hi() {

}
