package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type SessionService interface {
	GetUserBy(ctx context.Context, sessionID string) (entity.User, error)
	Generate(ctx context.Context, user entity.User) (entity.Session, error)
}
