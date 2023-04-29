package server

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
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

	box := rice.MustFindBox("../html")
	r.StaticFS("/html", http.FileSystem(box.HTTPBox()))

	return r
}
