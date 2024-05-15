package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/template"
	"github.com/abbasfisal/ecommerce-go/internal/client/contract"
	"github.com/abbasfisal/ecommerce-go/internal/client/requests"
	"github.com/abbasfisal/ecommerce-go/internal/pkg"
	PublicSrv "github.com/abbasfisal/ecommerce-go/internal/public/contract"
	"github.com/abbasfisal/ecommerce-go/internal/public/service"
	sessionContract "github.com/abbasfisal/ecommerce-go/internal/session/contract"
	"github.com/abbasfisal/ecommerce-go/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Client struct {
	authSrv    contract.AuthService
	sessionSrv sessionContract.SessionService
	publicSrv  PublicSrv.PublicSrv
}

func NewClient(clientAuthSrv contract.AuthService, sessionSrv sessionContract.SessionService, publicSrv PublicSrv.PublicSrv) Client {
	return Client{
		authSrv:    clientAuthSrv,
		sessionSrv: sessionSrv,
		publicSrv:  publicSrv,
	}
}

func (h Client) ShowRegisterForm(c *gin.Context) {
	//check cookie if not exist then show it
	if h.checkClientCookie(c) {
		return
	}

	//pass phase number
	c.HTML(http.StatusOK, "user-register.html", nil)
	return
}

func (h Client) checkClientCookie(c *gin.Context) bool {
	sessionID, sessErr := c.Cookie("client-session")
	if sessErr == nil {
		//session exists
		fmt.Println("session exist")
		user, err := h.sessionSrv.GetUserBy(context.TODO(), sessionID)
		if err == nil && user.ID > 0 {

			page, errPage := strconv.Atoi(c.Query("page"))
			if errPage != nil || page < 1 {
				page = 1
			}

			products, totalCount, err := h.publicSrv.GetProducts(context.TODO(), page)

			if totalCount == 0 {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					fmt.Println("\n show all products internal error ", err)
					c.HTML(http.StatusInternalServerError, "500.html", nil)
					return true
				}

				fmt.Println("\n show all products , no records was exists  ", err)
				c.HTML(http.StatusOK, "index.html", template.Data{
					Message:    "there is no products in table",
					StatusCode: 404,
				})

				return true
			}

			fmt.Println("here ", sessionID)
			c.HTML(http.StatusPermanentRedirect, "index.html", template.Data{
				Data: gin.H{
					"Products":    products,
					"HasPrev":     page > 1,
					"PrevPage":    page - 1,
					"Pages":       util.GeneratePageNumbers(page, int(totalCount)),
					"HasNext":     len(products) == service.PerPage,
					"NextPage":    page + 1,
					"CurrentPage": page,

					"User": user,
				},
			})

			return true
		} else {
			//delete session
			fmt.Println("sesson not found , delete session")
			c.SetCookie("client-session", "", -1, "/", "", false, true)
		}
	}
	return false
}

func (h Client) PostRegister(c *gin.Context) {

	if h.checkClientCookie(c) {
		return
	}

	_ = c.Request.ParseForm()
	for key, val := range c.Request.PostForm {
		fmt.Println(" key : ", key, " | value : ", val, " \n ")
	}

	var req requests.CreateRegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusBadRequest, "user-register.html", template.Data{
			Error: "invalid request parameters",
		})
		return
	}
	fmt.Println(req)

	//-----validate
	validate := pkg.NewValidate()
	if err := validate.Struct(&req); err != nil {
		validationError := pkg.CollectAndTranslateValidationErrors(err)

		c.HTML(http.StatusTemporaryRedirect, "user-register.html", template.Data{
			Error:           "validation failed",
			ValidationError: validationError,
			OldData:         util.StructToMap(req),
		})
		return
	}

	//check uniqueness of mobile and nic
	if ok := h.authSrv.CheckUniquePhoneAndNIC(context.TODO(), req.Mobile, req.NationalCode); !ok {
		c.HTML(http.StatusBadRequest, "user-register.html", template.Data{
			Error: "phone number Or national code is already exists",
			ValidationError: []map[string]any{
				{"Mobile|NIC": "شماره موبایل یا کد ملی تکراری است"},
			},
			OldData: util.StructToMap(req),
		})
		return
	}

	//--- end validation
	//call service // call repo // store
	user, err := h.authSrv.Register(context.TODO(), req)
	if err != nil {
		fmt.Println("error register a new user  ", err)
		c.HTML(http.StatusInternalServerError, "500.html", template.Data{
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
			"User": user,
		},
	})
	return

}

func (h Client) ShowLoginForm(c *gin.Context) {

	fmt.Println("show Login hit")

	if h.checkClientCookie(c) {
		return
	}

	c.HTML(http.StatusOK, "user-login.html", nil)
	return
}

func (h Client) PostLoginForm(c *gin.Context) {

	fmt.Println("Post Login hit")

	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println("\n\t parse form error : ", err)
		return
	}

	var req requests.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("\n \t bind error ", err)
		c.HTML(http.StatusBadRequest, "login.html", template.Data{
			Error: "binding data failed",
		})
		return
	}

	fmt.Println("request", req)

	if h.checkClientCookie(c) {
		return
	}

	user, loginErr := h.authSrv.Login(context.TODO(), req)
	fmt.Println("before ", user, loginErr)
	if loginErr != nil {
		fmt.Println("her  ess")
		if errors.Is(loginErr, gorm.ErrRecordNotFound) {
			fmt.Println("inseide")
			fmt.Println("record not found", loginErr)
			c.HTML(http.StatusNotFound, "user-login.html", template.Data{
				Error:   loginErr.Error(),
				Message: "record not found",
			})
			return
		} else {
			fmt.Println("internal server err", loginErr)
			c.HTML(http.StatusInternalServerError, "500.html", nil)
			return
		}
	}

	fmt.Println("user found successfully", user)

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
