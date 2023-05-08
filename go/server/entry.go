package server

import (
	"golib/server/container"
	"golib/server/report"
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
	container.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// All test cases go here.
	r.GET("/api/report/:id", report.Detail)
	r.GET("/api/report", report.List)
	r.POST("/api/report", report.Create)

	r.StaticFS("/html", gin.Dir("html", true))

	return r
}
