package routes

import (
	"github.com/abbasfisal/ecommerce-go/cmd/http_server/handler"
	"github.com/abbasfisal/ecommerce-go/repository"
	"github.com/abbasfisal/ecommerce-go/services"
	"github.com/gin-gonic/gin"
)

func (s ServerApis) SetAdminRoutes(r *gin.RouterGroup) {
	// routes url => v1/admin
	grp := r.Group("/admin")

	repo := repository.NewAdmin(s.Db)
	srv := services.NewAdminService(repo)
	hnd := handler.NewAdminHandler(srv)

	grp.GET("/login", hnd.ShowLogin)
	grp.POST("/login", hnd.Login)
}
