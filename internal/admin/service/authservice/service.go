package authservice

import (
	"context"
	"errors"
	"github.com/abbasfisal/ecommerce-go/internal/admin/contract"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo contract.AuthRepo
}

func New(repo contract.AuthRepo) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (s AuthService) Login(ctx context.Context, request requests.LoginRequest) (entity.User, error) {

	//check user by mobile and type = admin
	user, err := s.repo.GetUserBy(request.PhoneNumber)
	if err != nil {
		return entity.User{}, err
	}
	//compare hashed password
	hashCheckErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if hashCheckErr != nil {
		return entity.User{}, errors.New("record not found")
	}
	return user, nil
}
