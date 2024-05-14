package routes

import (
	"github.com/abbasfisal/ecommerce-go/internal/client/handlers"
	ClientAuthRepo "github.com/abbasfisal/ecommerce-go/internal/client/repository"
	ClientAuthSrv "github.com/abbasfisal/ecommerce-go/internal/client/service"
	"github.com/abbasfisal/ecommerce-go/internal/database/mysql"
	sessionRepoResolver "github.com/abbasfisal/ecommerce-go/internal/session/repository"
	sessionSrvResolver "github.com/abbasfisal/ecommerce-go/internal/session/service"
	"github.com/gin-gonic/gin"
)

func (s ServerApis) setClientRoutes(r *gin.Engine) {

	//init client handler
	authRepo := ClientAuthRepo.NewAuthenticateRepository(mysql.Get())
	authSrv := ClientAuthSrv.NewAuthenticateService(authRepo)

	sessionRepo := sessionRepoResolver.NewSessionRepository(mysql.Get())
	sessionSrv := sessionSrvResolver.NewSessionService(sessionRepo)

	hdl := handlers.NewClient(authSrv, sessionSrv)

	r.GET("/register", hdl.ShowRegisterForm)
	r.POST("/register", hdl.PostRegister)

	r.GET("/login", hdl.ShowLoginForm)
	r.POST("/login", hdl.PostLoginForm)

}
