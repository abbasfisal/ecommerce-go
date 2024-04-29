package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type AuthService interface {
	Login(ctx context.Context, request requests.LoginRequest) (entity.User, error)
}
