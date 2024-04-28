package authservice

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/contract"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
)

type AuthService struct {
	repo contract.AuthRepo
}

func New(repo contract.AuthRepo) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (s AuthService) Login(ctx context.Context, request requests.LoginRequest) {
}
