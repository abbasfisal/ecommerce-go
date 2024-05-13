package handlers

import (
	"context"
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/template"
	"github.com/abbasfisal/ecommerce-go/internal/client/contract"
	"github.com/abbasfisal/ecommerce-go/internal/client/requests"
	sessionContract "github.com/abbasfisal/ecommerce-go/internal/session/contract"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Client struct {
	authSrv    contract.AuthService
	sessionSrv sessionContract.SessionService
}

func NewClient(clientAuthSrv contract.AuthService, sessionSrv sessionContract.SessionService) Client {
	return Client{
		authSrv:    clientAuthSrv,
		sessionSrv: sessionSrv,
	}
}

func (h Client) ShowRegisterForm(c *gin.Context) {
	//check cookie if not exist then show it

	//pass phase number
	c.HTML(http.StatusOK, "user-register.html", nil)
	return
}

func (h Client) PostRegister(c *gin.Context) {
	fmt.Println("hi")
	sessionID, sessErr := c.Cookie("client-session")
	if sessErr == nil {
		//session exists
		user, err := h.sessionSrv.GetUserBy(context.TODO(), sessionID)
		if err == nil {
			fmt.Println("here ", sessionID)
			c.HTML(http.StatusPermanentRedirect, "index.html", template.Data{
				Data: gin.H{
					"user": user,
				},
			})
			return
		} else {
			//delete session
			c.SetCookie("client-session", "", 0, "/", "", false, true)
		}
	}

	fmt.Println("here33")
	for key, val := range c.Request.PostForm {
		fmt.Println(" key : ", key, " | value : ", val, " \n ")
	}
	fmt.Println("here 444")
	//check cookie exist? ->redirect (index)

	var req requests.CreateRegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusBadRequest, "user-register.html", template.Data{
			Error: "invalid request parameters",
		})
		return
	}
	fmt.Println(req)

	//validate

	//check uniqueness of mobile and nic
	if ok := h.authSrv.CheckUniquePhoneAndNIC(context.TODO(), req.Mobile, req.NationalCode); !ok {
		c.HTML(http.StatusBadRequest, "user-register.html", template.Data{
			Error: "phone number Or national code is already exists",
		})
		return
	}

	//call service // call repo // store
	user, err := h.authSrv.Register(context.TODO(), req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "user-register.html", template.Data{
			Error:      "Failed to register a client ",
			Meta:       gin.H{"error_text": err},
			StatusCode: 500,
		})
		return
	}

	fmt.Println("\n user------- ", user)
	//generate session and store it in db
	session, SessErr := h.sessionSrv.Generate(context.TODO(), user)
	if SessErr != nil {
		//redirect to login html
		fmt.Println("error while generating session", err)
		c.HTML(http.StatusPermanentRedirect, "user-login", nil)
		return
	}

	fmt.Println("user created successfully")
	//generate cookie
	c.SetCookie("client-session", session.SessionID, 3600, "/", "", false, true)
	c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{
		"msg": "user loggedId",
		"data": gin.H{
			"user": user,
		},
	})
	return

}
