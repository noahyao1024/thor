package server

import (
	ca "golib/server/cases"
	"golib/server/container"
	"golib/server/model"
	"golib/server/report"
	"net/http"
	"time"

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
	r := stage1()
	r = stage2(r)
	return r
}

func stage1() *gin.Engine {
	container.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		db := container.GetLocalDB()

		if db == nil {
			c.JSON(http.StatusOK, gin.H{
				"path":    container.GetSQLLitePath(),
				"message": "pong",
				"error":   "db is nil",
			})
			return
		}

		err := db.Exec(`CREATE TABLE IF NOT EXISTS reports (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			data TEXT NULL,
			status TEXT NULL,
			create_time DATETIME DEFAULT (datetime('now')),
			modify_time DATETIME DEFAULT (datetime('now', 'localtime'))
		  );
		  CREATE TABLE IF NOT EXISTS cases (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			template_id INTEGER,
			template_name TEXT NULL,
			report_id INTEGER NULL,
			status TEXT NULL,
			data TEXT NULL,
			apis TEXT NULL,
			create_time DATETIME DEFAULT (datetime('now')),
			modify_time DATETIME DEFAULT (datetime('now', 'localtime'))
		  );
		  `)

		report := &model.Report{Name: "Hello", Data: "World"}
		report.Status = "init"
		report.CreateTime = time.Now()
		db.Create(&report)

		reports := make([]*model.Report, 0)
		db.Where(&model.Report{}).Limit(100).Order("id DESC").Find(&reports)

		c.JSON(http.StatusOK, gin.H{
			"path":    container.GetSQLLitePath(),
			"message": "pong",
			"err":     err.Error,
			"reports": reports,
		})
	})

	// All test cases go here.
	r.GET("/api/report/:id", report.Detail)
	r.GET("/api/report", report.List)
	r.POST("/api/report", report.Create)

	r.POST("/api/case/execute/:report_id", ca.BatchRun)

	r.StaticFS("/html", gin.Dir("html", true))

	return r
}
