package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/template"
	"github.com/abbasfisal/ecommerce-go/internal/public/contract"
	"github.com/abbasfisal/ecommerce-go/internal/public/service"
	sessionContract "github.com/abbasfisal/ecommerce-go/internal/session/contract"
	"github.com/abbasfisal/ecommerce-go/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Public struct {
	PublicSrv  contract.PublicSrv
	sessionSrv sessionContract.SessionService
}

func NewPublic(publicSrv contract.PublicSrv, sessionSrv sessionContract.SessionService) Public {
	return Public{
		PublicSrv:  publicSrv,
		sessionSrv: sessionSrv,
	}
}

func (h Public) ShowIndexSite(c *gin.Context) {
	fmt.Println("show index page hit")

	//todo: get perPage query
	page, errPage := strconv.Atoi(c.Query("page"))
	if errPage != nil || page < 1 {
		page = 1
	}

	products, totalCount, err := h.PublicSrv.GetProducts(context.TODO(), page)
	fmt.Println("\n products ", products, totalCount, err)
	if totalCount == 0 {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			//
			fmt.Println("\n show all products internal error ", err)
			c.HTML(http.StatusInternalServerError, "500.html", nil)
			return
		}

		fmt.Println("\n show all products , no records was exists  ", err)
		c.HTML(http.StatusOK, "index.html", template.Data{
			Message:    "there is no products in table",
			StatusCode: 404,
		})
		return
	}

	fmt.Println("\n ---- here end")
	c.HTML(http.StatusOK, "index.html", template.Data{
		Data: gin.H{
			"Products":    products,
			"HasPrev":     page > 1,
			"PrevPage":    page - 1,
			"Pages":       util.GeneratePageNumbers(page, int(totalCount)),
			"HasNext":     len(products) == service.PerPage,
			"NextPage":    page + 1,
			"CurrentPage": page,
		},
	})
	return
}
