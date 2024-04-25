package handler

import (
	"fmt"
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

func (h AdminHandler) ShowLogin(c *gin.Context) {
	fmt.Println("show login form")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"name": "ali",
	})
}

type Req struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (h AdminHandler) Login(c *gin.Context) {

	fmt.Println("post login form")
	var r Req
	err := c.ShouldBind(&r)
	fmt.Println(r)

	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"Err": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"message": "you successfully registered"})

}
