package cases

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Wrap ...
func Wrap(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
