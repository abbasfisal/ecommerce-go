package routes

import (
	"github.com/abbasfisal/ecommerce-go/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerApis struct {
	Db     *gorm.DB
	Config config.Config
	//logs

}

func NewServerApis(db *gorm.DB, config *config.Config) ServerApis {
	return ServerApis{
		Db:     db,
		Config: *config,
	}
}

func (s ServerApis) Run() {

	engin := gin.New()

	SetHealthRoutes(engin.Group("/health"))

	group := engin.Group("/v1")

	s.SetAdminRoutes(group)

	engin.Run(":8080")
}
