package service

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/client/contract"
	"github.com/abbasfisal/ecommerce-go/internal/client/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type AuthenticateService struct {
	Repo contract.AuthRepo
}

func NewAuthenticateService(
	repo contract.AuthRepo,
) AuthenticateService {

	return AuthenticateService{
		Repo: repo,
	}
}

func (a AuthenticateService) CheckUniquePhoneAndNIC(ctx context.Context, phone string, nic string) bool {
	return a.Repo.IsUniquePhoneAndNIC(ctx, phone, nic)
}

func (a AuthenticateService) Register(ctx context.Context, req requests.CreateRegisterRequest) (entity.User, error) {
	return a.Repo.CreateNewClient(ctx, req)
}

func (a AuthenticateService) Login(ctx context.Context, req requests.LoginRequest) (entity.User, error) {
	return a.Repo.GetUserBy(ctx, req)
}
