package repository

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"gorm.io/gorm"
)

type Authenticate struct {
	db *gorm.DB
}

func NewAuthenticate(db *gorm.DB) Authenticate {
	return Authenticate{
		db: db,
	}
}

func (a Authenticate) List(c context.Context) error {
	return nil
}

func (a Authenticate) CheckAdminExists(c context.Context, loginReq requests.LoginRequest) (entity.User, error) {
	return entity.User{}, nil
}
