package handler

import (
	"github.com/abbasfisal/ecommerce-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminHandler struct {
	srv services.AdminService // u can use interfaces for admin service
}

func NewAdminHandler(srv services.AdminService) AdminHandler {
	return AdminHandler{
		srv: srv,
	}
}

func (h AdminHandler) Login(c *gin.Context) {
	//validation

	h.srv.List(c)
	c.JSON(http.StatusOK, gin.H{"message ": "this is login"})
}
