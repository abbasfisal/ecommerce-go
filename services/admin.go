package services

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/contract"
)

type AdminService struct {
	repo contract.AdminContract
}

func NewAdminService(repo contract.AdminContract) AdminService {
	return AdminService{
		repo: repo,
	}
}
func (s AdminService) List(c context.Context) {
	err := s.repo.List(c)
	if err != nil {
		return
	}
}
