package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/client/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"gorm.io/gorm"
)

type AuthenticateRepository struct {
	Db *gorm.DB
}

func NewAuthenticateRepository(db *gorm.DB) AuthenticateRepository {

	return AuthenticateRepository{
		Db: db,
	}
}
func (a AuthenticateRepository) IsUniquePhoneAndNIC(ctx context.Context, phone string, nic string) bool {
	var count int64
	a.Db.Where("phone_number = ? or national_code = ? ", phone, nic).Count(&count)

	return count == 0
}

func (a AuthenticateRepository) CreateNewClient(ctx context.Context, req requests.CreateRegisterRequest) (entity.User, error) {

	user := entity.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.Mobile,
		Password:     req.Password,
		Type:         "client",
		NationalCode: req.NationalCode,
		Address: entity.Address{
			Floor:       req.Floor,
			Phase:       "a",
			Number:      req.Number,
			Block:       "block2",
			Description: req.Description,
		},
	}
	fmt.Println("user data to insert : ", user)

	if err := a.Db.Create(&user).Error; err != nil {
		fmt.Println("error : ", err)
		return entity.User{}, errors.New("create a new user was failed")
	}

	return user, nil
}

func (a AuthenticateRepository) GetUserBy(ctx context.Context, req requests.LoginRequest) (entity.User, error) {
	//todo: encrypt password
	var user entity.User

	if err := a.Db.Where("phone_number=? and password = ? ", req.Mobile, req.Password).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
