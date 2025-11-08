package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WhoamiHandler(c *gin.Context) {
	user, exists := c.Get("user")

	if exists {
		c.JSON(http.StatusOK, gin.H{"message": "whoami", "user": user})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}
}
