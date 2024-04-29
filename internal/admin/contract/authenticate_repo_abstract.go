package contract

import (
	"context"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
)

type AuthRepo interface {
	List(c context.Context) error
	CheckAdminExists(c context.Context, loginReq requests.LoginRequest) (entity.User, error)
	GetUserBy(phoneNumber string) (entity.User, error)
}
