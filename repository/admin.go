package repository

import (
	"context"
	"gorm.io/gorm"
)

type Admin struct {
	db *gorm.DB
}

func NewAdmin(db *gorm.DB) Admin {
	return Admin{
		db: db,
	}
}

func (a Admin) List(c context.Context) error {
	return nil
}
