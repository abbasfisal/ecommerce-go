package handler

import (
	"fmt"
	"github.com/abbasfisal/ecommerce-go/services"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
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

	// check session existence and user type

	// else show login form

	tmpl, err := template.ParseFiles("template/admin/login.html")
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	type Data struct {
		Name string
	}

	mydata := Data{Name: "reza"}
	err2 := tmpl.Execute(c.Writer, mydata)
	if err2 != nil {
		log.Fatal(err2)
	}

}
