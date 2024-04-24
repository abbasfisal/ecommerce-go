package contract

import (
	"context"
)

//for repository contract

type AdminContract interface {
	List(c context.Context) error
}
