package util

import (
	"github.com/abbasfisal/ecommerce-go/internal/admin/transport/http/template"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error500(ctx *gin.Context) {
	ctx.HTML(http.StatusInternalServerError, "500.html", template.Data{
		Message:    "",
		Error:      "",
		StatusCode: 500,
	})

}
