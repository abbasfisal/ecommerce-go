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
	productSrv  contract.ProductService
}

func NewAdminHandler(
	authSrv contract.AuthService,
	sessionSrv sessionContract.SessionService,
	categorySrv contract.CategoryService,
	productSrv contract.ProductService,
) AdminHandler {
	return AdminHandler{
		authSrv:     authSrv,
		sessionSrv:  sessionSrv,
		categorySrv: categorySrv,
		productSrv:  productSrv,
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
	//todo: validation for image

	imageFile, err := c.FormFile("image")
	if err != nil {
		c.HTML(http.StatusBadRequest, "create-category.html", template.Data{
			Message:    "get image failed",
			Error:      err.Error(),
			StatusCode: 0,
		})
		return
	}
	fmt.Println("\n--image header", imageFile.Header)

	imageMustStorePath := util.GenerateFilename(imageFile.Filename)
	fmt.Println("\n--imageMustStorePath", imageMustStorePath)

	//todo: read upload main dir from env files
	if err := c.SaveUploadedFile(imageFile, "media/categories/"+imageMustStorePath); err != nil {
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
		os.Remove(imageMustStorePath)

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
	})
	return
}
func (h AdminHandler) ShowCategoryList(c *gin.Context) {
	categories, err := h.categorySrv.GetAll(c)
	if err != nil {
		util.Error500(c)
		return
	}

	c.HTML(http.StatusOK, "list-category.html", template.Data{
		Message: "",
		Data: map[string]any{
			"Categories": categories,
		},
	})
	return
}
func (h AdminHandler) ShowCreateProduct(c *gin.Context) {

	var tData template.Data
	// get all categories
	categories, err := h.categorySrv.GetAll(c)
	if err != nil {
		tData.Message = "create a new category"
		tData.StatusCode = 404

	} else {
		tData.StatusCode = 200
		tData.Data = map[string]any{"categories": categories}
	}
	c.HTML(http.StatusOK, "create-product.html", tData)
	return
}

func (h AdminHandler) StoreProduct(c *gin.Context) {
	fmt.Println("\n\t --- store product hit")
	categories, _ := h.categorySrv.GetAll(c)

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("\n\t -- multipart err : ", err)
		util.Error500(c)
		return
	}

	var req requests.CreateProductRequest
	//delete below
	for k, v := range c.Request.PostForm {
		fmt.Println("key: ", k, "\t value : ", v)
	}

	//shouldBind
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("\n\t --- bind error ", err)
		c.HTML(http.StatusBadRequest, "create-product.html", template.Data{
			Message:    "bind post form goes wrong",
			Error:      err.Error(),
			StatusCode: 429,
			Data: map[string]any{
				"categories": categories,
			},
		})
		return
	}
	//validation
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&req); err != nil {
		c.HTML(http.StatusBadRequest, "create-product.html", template.Data{
			Message:    "validation error ",
			Error:      err.Error(),
			StatusCode: 429,
			Data: map[string]any{
				"categories": categories,
			},
		})
		return
	}
	fmt.Println("\n\t--- request bind  ", req)

	images := form.File["images"]
	var imagesPath []string
	for index, image := range images {
		//todo: validate image
		fmt.Println("\n\t - a ", index, "\n\t -- ", image.Filename)

		imageUploadPath := util.GenerateFilename(image.Filename)
		if err := c.SaveUploadedFile(image, "media/products/"+imageUploadPath); err != nil {
			fmt.Println("\n--- internal server error | storing uploaded image ", err)
			util.Error500(c)
			return
		}
		imagesPath = append(imagesPath, imageUploadPath)
	}

	//use product service
	product, pErr := h.productSrv.Create(c, req, imagesPath)
	if pErr != nil {
		fmt.Println("\n\t --- failed create a new Product", pErr)
		util.Error500(c)
		return
	}

	fmt.Println("\n\t --- product created successfully ", product)
	c.HTML(http.StatusCreated, "create-product.html", template.Data{
		Message:    "product created successfully",
		StatusCode: 201,
		Data:       map[string]any{"product": product, "categories": categories},
	})
	return

}

func (h AdminHandler) ShowProductList(c *gin.Context) {
	fmt.Println("\n show product list hit")
	products, err := h.productSrv.List(c)
	fmt.Println("\n\t======== products ", products)
	if err != nil {
		c.HTML(http.StatusOK, "list-product.html", template.Data{
			Message:    "record not found",
			Error:      err.Error(),
			StatusCode: 404,
		})
		return
	}

	c.HTML(http.StatusOK, "list-product.html", template.Data{
		Message:    "successfully get products",
		StatusCode: 200,
		Data: map[string]any{
			"Products": products,
		},
	})
	return

}
