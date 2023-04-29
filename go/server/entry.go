package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListenAtBackground ...
func ListenAtBackground(port string) {
	r := listen()
	go r.Run(":" + port)
}

// Listen ...
func Listen(port string) {
	r := listen()
	r.Run(":" + port)
}

func listen() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.StaticFS("/html", gin.Dir("html", true))

	return r
}
