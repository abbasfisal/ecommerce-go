package middleware

import (
	"fmt"
	"github.com/abbasfisal/ecommerce-go/internal/database/mysql"
	"github.com/abbasfisal/ecommerce-go/internal/session/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAdmin(c *gin.Context) {
	fmt.Println("\n--- is admin middleware fired -----")

	sessionID, err := c.Cookie("session-admin")
	if err != nil {
		c.Redirect(http.StatusPermanentRedirect, "/login-admin")
		c.Abort()
		return
	}

	sessionRepo := repository.NewSessionRepository(mysql.Get())

	user, sessErr := sessionRepo.GetUserBySession(c, sessionID)
	fmt.Println("\n--- IsAdmin middleware User data --", user)
	fmt.Println("\n--- session error ---", sessErr)
	if sessErr != nil || user.Type != "admin" {
		fmt.Println("\n---- user not found | user.type is not admin ")
		c.Redirect(http.StatusPermanentRedirect, "/login-admin")
		c.Abort()
		return
	}

	fmt.Println("\n---- user logged int by cookie and was admin ---")

	c.Next()
}
