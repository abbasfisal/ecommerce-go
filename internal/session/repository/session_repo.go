package repository

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"gorm.io/gorm"
)

type SessionRepository struct {
	Db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return SessionRepository{
		Db: db,
	}
}
func (r SessionRepository) GetUserBySession(ctx context.Context, sessionID string) (entity.User, error) {
	var u entity.User
	r.Db.Where("session_id=?", sessionID).Model(&u)

	return u, nil
}

func (r SessionRepository) Generate(ctx context.Context, session entity.Session) (entity.Session, error) {
	result := r.Db.Create(session)

	if result.RowsAffected == 0 {
		return entity.Session{}, result.Error
	}

	return session, nil
}
