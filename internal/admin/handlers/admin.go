package handlers

import (
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/admin/contract"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/entity"
	sessionContract "github.com/abbasfisal/ecommerce-go/internal/session/contract"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type AdminHandler struct {
	authSrv    contract.AuthService
	sessionSrv sessionContract.SessionService
}

func NewAdminHandler(authSrv contract.AuthService, sessionSrv sessionContract.SessionService) AdminHandler {
	return AdminHandler{
		authSrv:    authSrv,
		sessionSrv: sessionSrv,
	}
}

func (h AdminHandler) ShowLogin(c *gin.Context) {
	fmt.Println("show login form")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"name": "ali",
	})

}

func (h AdminHandler) Login(c *gin.Context) {
	fmt.Println("post login form")

	var req requests.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Println("binding error : ", err)
		c.HTML(http.StatusBadRequest, "500.html", nil)
		return
	}
	fmt.Println(req)

	//validation
	validate := validator.New()
	valErr := validate.Struct(req)
	if valErr != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"validationError": valErr,
		})
		return
	}

	//service
	//adminCheck, checkErr := h.srv.CheckAdminExists(c, req)
	//if checkErr != nil {
	//	c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "username | password was incorrect"})
	//	c.Abort()
	//	return
	//}
	//
	////compare hashed password
	//hashErr := bcrypt.CompareHashAndPassword([]byte(adminCheck.Password), []byte(req.Password))
	//if hashErr != nil {
	//	c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "username | password was incorrect"})
	//	c.Abort()
	//	return
	//}

	//generate session and store in session table
	var session entity.Session
	session.SessionID = uuid.New().String()
	//session.UserID = adminCheck.ID
	session.CreatedAt = time.Now()
	session.UpdatedAt = session.CreatedAt

	c.SetCookie("session", session.SessionID, 3600, "/", "", true, true)
	//sessionId = uuid.New()

	c.Redirect(http.StatusOK, "/dashboard")
	//c.HTML(http.StatusOK, "dashboard.html", gin.H{"message": "you successfully registered"})

}
