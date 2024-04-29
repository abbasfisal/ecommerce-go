package routes

import (
	"github.com/abbasfisal/ecommerce-go/internal/admin/handlers"
	"github.com/abbasfisal/ecommerce-go/internal/admin/repository"
	adminAuthSrv "github.com/abbasfisal/ecommerce-go/internal/admin/service/authservice"
	"github.com/abbasfisal/ecommerce-go/internal/middleware"
	sessionRepoResolver "github.com/abbasfisal/ecommerce-go/internal/session/repository"
	sessionSrvResolver "github.com/abbasfisal/ecommerce-go/internal/session/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s ServerApis) SetAdminRoutes(r *gin.Engine) {

	authRepo := repository.NewAuthenticate(s.Db)
	authSvc := adminAuthSrv.New(authRepo)

	sessionRepo := sessionRepoResolver.NewSessionRepository(s.Db)
	sessionSrv := sessionSrvResolver.NewSessionService(sessionRepo)

	hnd := handlers.NewAdminHandler(authSvc, sessionSrv)

	//=>login-admin
	r.GET("/login-admin", hnd.ShowLogin)
	r.POST("/login-admin", hnd.Login)

	grp := r.Group("v1/admin", middleware.IsAdmin)
	{
		//group route =>v1/admin/dashboard
		grp.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard.html", nil)
			return
		})

	}
}
