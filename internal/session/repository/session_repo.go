package repository

import (
	"context"
	"errors"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"github.com/google/uuid"
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
	//todo: put user data like type and name into sessions table
	var u entity.User
	result := r.Db.Table("users").Select("users.*, sessions.*").
		Joins("LEFT JOIN sessions ON users.id = sessions.user_id").
		Where("sessions.session_id = ?", sessionID).
		Find(&u)
	if result.RowsAffected <= 0 {
		return u, result.Error
	}
	return u, nil
}

func (r SessionRepository) Generate(ctx context.Context, user entity.User) (entity.Session, error) {
	session := entity.Session{
		SessionID: uuid.New().String(),
		UserID:    user.ID,
	}

	result := r.Db.Create(&session)

	if result.RowsAffected == 0 {
		return entity.Session{}, errors.New("can not create session in session table")
	}

	return session, nil
}
