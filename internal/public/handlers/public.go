package handlers

import (
	"github.com/abbasfisal/ecommerce-go/internal/public/contract"
	sessionContract "github.com/abbasfisal/ecommerce-go/internal/session/contract"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.HTML(http.StatusOK, "index.html", nil)
	return
}
