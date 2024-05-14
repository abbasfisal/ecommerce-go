package routes

import (
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/client/handlers"
	ClientAuthRepo "github.com/abbasfisal/ecommerce-go/internal/client/repository"
	ClientAuthSrv "github.com/abbasfisal/ecommerce-go/internal/client/service"
	"github.com/abbasfisal/ecommerce-go/internal/database/mysql"
	sessionRepoResolver "github.com/abbasfisal/ecommerce-go/internal/session/repository"
	sessionSrvResolver "github.com/abbasfisal/ecommerce-go/internal/session/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

	r.GET("/hi", func(c *gin.Context) {
		Hi(c)
		return
		fmt.Println("after hi func")
	})
}

func Hi(c *gin.Context) {
	c.HTML(http.StatusPermanentRedirect, "index.html", nil)
	return
}
