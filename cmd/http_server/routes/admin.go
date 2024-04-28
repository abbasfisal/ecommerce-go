package routes

import (
	"github.com/abbasfisal/ecommerce-go/internal/admin/handlers"
	"github.com/abbasfisal/ecommerce-go/internal/admin/repository"
	adminAuthSrv "github.com/abbasfisal/ecommerce-go/internal/admin/service/authservice"
	sessionRepoResolver "github.com/abbasfisal/ecommerce-go/internal/session/repository"
	sessionSrvResolver "github.com/abbasfisal/ecommerce-go/internal/session/service"
	"github.com/gin-gonic/gin"
)

func (s ServerApis) SetAdminRoutes(r *gin.RouterGroup) {
	// routes url => v1/admin
	grp := r.Group("/admin")

	authRepo := repository.NewAuthenticate(s.Db)
	authSvc := adminAuthSrv.New(authRepo)

	sessionRepo := sessionRepoResolver.NewSessionRepository(s.Db)
	sessionSrv := sessionSrvResolver.NewSessionService(sessionRepo)

	hnd := handlers.NewAdminHandler(authSvc, sessionSrv)

	grp.GET("/login", hnd.ShowLogin)
	grp.POST("/login", hnd.Login)
}
