package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetHealthRoutes(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}
