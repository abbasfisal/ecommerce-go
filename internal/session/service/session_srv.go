package service

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"github.com/abbasfisal/ecommerce-go/internal/session/contract"
)

type SessionService struct {
	repo contract.SessionRepository
}

func NewSessionService(repo contract.SessionRepository) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

func (s SessionService) GetUserBy(ctx context.Context, sessionID string) (entity.User, error) {
	return s.repo.GetUserBySession(ctx, sessionID)
}

func (s SessionService) Generate(ctx context.Context, user entity.User) (entity.Session, error) {
	return s.repo.Generate(ctx, user)
}
