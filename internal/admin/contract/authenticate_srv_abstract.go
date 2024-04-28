package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
)

type AuthService interface {
	Login(ctx context.Context, request requests.LoginRequest)
}
