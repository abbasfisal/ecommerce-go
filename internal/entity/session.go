package entity

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	SessionID string
	UserID    uint //adminID
}
