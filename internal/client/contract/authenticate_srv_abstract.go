package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/client/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type AuthService interface {
	CheckUniquePhoneAndNIC(ctx context.Context, phone string, nic string) bool
	Register(ctx context.Context, req requests.CreateRegisterRequest) (entity.User, error)
	Login(ctx context.Context, req requests.LoginRequest) (entity.User, error)
}
