package handlers

import (
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/admin/contract"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/requests"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/template"
	sessionContract "github.com/abbasfisal/ecommerce-go/internal/session/contract"
	"github.com/abbasfisal/ecommerce-go/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"os"
)

type AdminHandler struct {
	authSrv     contract.AuthService
	sessionSrv  sessionContract.SessionService
	categorySrv contract.CategoryService
}

func NewAdminHandler(authSrv contract.AuthService, sessionSrv sessionContract.SessionService, categorySrv contract.CategoryService) AdminHandler {
	return AdminHandler{
		authSrv:     authSrv,
		sessionSrv:  sessionSrv,
		categorySrv: categorySrv,
	}
}

func (h AdminHandler) ShowLogin(c *gin.Context) {
	fmt.Println("show login form")

	sessionID, err := c.Cookie("session-admin")
	if err == nil {
		//cookie is existed
		user, err := h.sessionSrv.GetUserBy(c, sessionID)
		if err != nil || user.Type != "admin" {

			c.SetCookie("session-admin", "", -1, "/", "", false, true)
			c.HTML(http.StatusOK, "login.html", nil)
			return
		}
		fmt.Println("\n -- session exist and was admin on show login form -- ")
		c.Redirect(http.StatusMovedPermanently, "v1/admin/dashboard")
		return
	}

	c.HTML(http.StatusOK, "login.html", nil)
	return
}

func (h AdminHandler) Login(c *gin.Context) {
	fmt.Println("\n---post login form")

	var req requests.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Println("binding error : ", err)
		c.HTML(http.StatusBadRequest, "500.html", nil)
		c.Abort()
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

	user, checkErr := h.authSrv.Login(c, req)
	if checkErr != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "username | password was incorrect",
			"meta":    checkErr,
		})
		c.Abort()
		return
	}

	sessionID, cookieErr := c.Cookie("session-admin")
	if cookieErr != nil {
		//cookie not exist
		session, sessErr := h.sessionSrv.Generate(c, user)
		if sessErr != nil {
			fmt.Println("\n --- generate session failed")
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"message": sessErr,
			})
			c.Abort()
			return
		}
		c.SetCookie("session-admin", session.SessionID, 3600, "/", "", false, true)

	} else {
		//cookie was existed
		user, err := h.sessionSrv.GetUserBy(c, sessionID)
		if err != nil || user.Type != "admin" {

			c.SetCookie("session-admin", "", -1, "/", "", false, true)
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"message": err,
			})
			c.Abort()
			return
		}
	}

	c.Redirect(http.StatusMovedPermanently, "v1/admin/dashboard")
	return
}

func (h AdminHandler) ShowCreateCategory(c *gin.Context) {
	fmt.Println("\n--- route categories triggered ---")
	c.HTML(http.StatusOK, "create-category.html", nil)
	return
}

func (h AdminHandler) StoreCategory(c *gin.Context) {
	fmt.Println("\n---upload store category", "\n\t form posted data ", c.Request.PostForm)
	var req requests.CreateCategoryRequest

	//bind
	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusBadRequest, "create-category.html", template.Data{
			Message: "validation failed",
			Error:   err.Error(),
		})

		fmt.Println("\n--- bind err : ", err)
		return
	}
	fmt.Println("\n\t --- request values", req, "\t", req.Status)

	//todo: use validator pkg

	imageFile, err := c.FormFile("image")
	if err != nil {
		c.HTML(http.StatusBadRequest, "create-category.html", template.Data{
			Message:         "get image failed",
			Error:           err.Error(),
			ValidationError: nil,
			StatusCode:      0,
			Data:            nil,
			Meta:            nil,
		})
		return
	}
	fmt.Println("\n--image header", imageFile.Header)

	imageMustStorePath := util.GenerateFilename(imageFile.Filename)
	fmt.Println("\n--imageMustStorePath", imageMustStorePath)

	if err := c.SaveUploadedFile(imageFile, imageMustStorePath); err != nil {
		fmt.Println("\n--- store image error :", template.Data{
			Message: "store image into disk failed",
			Error:   err.Error(),
		})
		return
	}
	fmt.Println("\n\t ---- req", &req, c)
	category, catErr := h.categorySrv.StoreCategory(c, &req, imageMustStorePath)

	if catErr != nil {
		fmt.Println("\n\t--- error from repository ", catErr)
		err := os.Remove(imageMustStorePath)
		if err != nil {
			log.Fatal(err)
			return
		}
		c.HTML(http.StatusBadRequest, "create-category.html", template.Data{
			Message:         "failed to create a record to db",
			Error:           err.Error(),
			ValidationError: nil,
			StatusCode:      0,
			Data:            nil,
			Meta:            nil,
		})

		return
	}

	c.HTML(http.StatusCreated, "create-category.html", template.Data{
		Message: "category created successfully",
		Data: map[string]any{
			"category": category,
		},
		Meta: nil,
	})
}
