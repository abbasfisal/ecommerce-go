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
	engin.LoadHTMLGlob("./template/*.html")

	engin.Static("/public", "./public")
	engin.Static("/media", "./media")

	SetHealthRoutes(engin.Group("/health"))

	//----------- admin routes
	s.SetAdminRoutes(engin)

	//----------- client routes
	s.setClientRoutes(engin)

	//----------- public routes
	s.SetPublicRoutes(engin)

	err := engin.Run(":" + s.Config.AppPort)
	if err != nil {
		log.Fatal("server error : ", err)
	}
}
