package routes

import (
	"github.com/abbasfisal/ecommerce-go/internal/admin/handlers"
	"github.com/abbasfisal/ecommerce-go/internal/admin/repository"
	adminAuthSrv "github.com/abbasfisal/ecommerce-go/internal/admin/service/authservice"
	"github.com/abbasfisal/ecommerce-go/internal/admin/service/categoryservice"
	"github.com/abbasfisal/ecommerce-go/internal/admin/service/productservice"
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

	categoryRepo := repository.NewCategoryRepository(s.Db)
	categorySrv := categoryservice.NewService(categoryRepo)

	productRepo := repository.NewProductRepository(s.Db)
	productSrv := productservice.NewService(productRepo)

	hnd := handlers.NewAdminHandler(authSvc, sessionSrv, categorySrv, productSrv)

	//=>login-admin
	r.GET("/login-admin", hnd.ShowLogin)
	r.POST("/login-admin", hnd.Login)

	grp := r.Group("v1/admin") //todo:add middleware IsAdmin
	{
		//group route =>v1/admin/dashboard
		grp.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard.html", nil)
			return
		})

		grp.GET("/categories", hnd.ShowCreateCategory)
		grp.POST("/categories", hnd.StoreCategory)

		grp.GET("/products", hnd.ShowCreateProduct)
		grp.POST("/products", hnd.StoreProduct)

	}
}
