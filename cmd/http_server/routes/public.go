package routes

import (
	"github.com/abbasfisal/ecommerce-go/internal/database/mysql"
	"github.com/abbasfisal/ecommerce-go/internal/public/handlers"
	"github.com/abbasfisal/ecommerce-go/internal/public/repository"
	"github.com/abbasfisal/ecommerce-go/internal/public/service"
	sessionRepoResolver "github.com/abbasfisal/ecommerce-go/internal/session/repository"
	sessionSrvResolver "github.com/abbasfisal/ecommerce-go/internal/session/service"
	"github.com/gin-gonic/gin"
)

func (s ServerApis) SetPublicRoutes(r *gin.Engine) {

	//session
	sessionRepo := sessionRepoResolver.NewSessionRepository(mysql.Get())
	sessionSrv := sessionSrvResolver.NewSessionService(sessionRepo)

	//repo
	publicRepo := repository.NewPublicRepository(mysql.Get())

	//service
	publicSrv := service.NewPublicService(publicRepo)

	//handler
	hdl := handlers.NewPublic(publicSrv, sessionSrv)

	r.GET("/", hdl.ShowIndexSite)
	r.GET("/products/:id", hdl.ShowByID)
}
