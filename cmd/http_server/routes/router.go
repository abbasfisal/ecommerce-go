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
	gin.SetMode(gin.DebugMode)
	engin.LoadHTMLGlob("./template/admin/*")

	engin.Static("/public", "./public")

	//todo : add route for 404 , 500

	SetHealthRoutes(engin.Group("/health"))

	s.SetAdminRoutes(engin)

	err := engin.Run(":" + s.Config.AppPort)
	if err != nil {
		log.Fatal("server error : ", err)
	}
}
