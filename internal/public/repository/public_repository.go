package repository

import "gorm.io/gorm"

type PublicRepository struct {
	Db *gorm.DB
}

func NewPublicRepository(db *gorm.DB) PublicRepository {
	return PublicRepository{
		Db: db,
	}
}
