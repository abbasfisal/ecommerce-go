package routes

import (
	"github.com/abbasfisal/ecommerce-go/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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
	engin.LoadHTMLGlob("./template/admin/*")
	//engin.LoadHTMLGlob("./template/*.html") //todo: try to load errors html

	engin.Static("/public", "./public")

	SetHealthRoutes(engin.Group("/health"))

	group := engin.Group("/v1")

	s.SetAdminRoutes(group)

	err := engin.Run(":" + s.Config.AppPort)
	if err != nil {
		log.Fatal("server error : ", err)
	}
}
