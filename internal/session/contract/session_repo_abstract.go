package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type SessionRepository interface {
	GetUserBySession(ctx context.Context, sessionID string) (entity.User, error)
	Generate(ctx context.Context, session entity.Session) (entity.Session, error)
}
