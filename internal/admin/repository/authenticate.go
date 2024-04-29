package repository

import (
	"context"
	"errors"
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

func (a Authenticate) GetUserBy(phoneNumber string) (entity.User, error) {
	//todo: change admin to const
	var user entity.User
	result := a.db.Where("phone_number=? and type='admin'", phoneNumber).First(&user)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("record not found")
	}
	return user, nil
}
