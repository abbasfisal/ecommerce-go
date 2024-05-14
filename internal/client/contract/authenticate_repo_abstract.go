package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/client/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type AuthRepo interface {
	IsUniquePhoneAndNIC(ctx context.Context, phone string, nic string) bool
	CreateNewClient(ctx context.Context, req requests.CreateRegisterRequest) (entity.User, error)
	GetUserBy(ctx context.Context, req requests.LoginRequest) (entity.User, error)
}
