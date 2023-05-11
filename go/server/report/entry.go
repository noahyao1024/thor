package report

import (
	"encoding/json"
	"fmt"
	"golib/server/container"
	"golib/server/model"
	"golib/server/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Detail ...
func Detail(c *gin.Context) {
	db := container.GetLocalDB()

	reportID := c.Param("id")

	report := &model.Report{}
	db.Where(&model.Report{}).Where("id = ?", reportID).Find(&report)
	if report.ID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    map[string]interface{}{},
		})
		return
	}

	cases := make([]*model.Case, 0)
	db.Where(&model.Case{}).Where("report_id = ?", reportID).Find(&cases)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": map[string]interface{}{
			"report": report,
			"cases":  cases,
		},
	})
}

// List ...
func List(c *gin.Context) {
	db := container.GetLocalDB()

	reports := make([]*model.Report, 0)
	db.Where(&model.Report{}).Limit(100).Order("id DESC").Find(&reports)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    reports,
	})
}

// Create ...
/*
curl 'localhost:9000/api/reports' -H 'Content-Type: application/json' -d '{
    "name": "Hello World",
    "data": "{\"SN\":\"PHBBBB\"}"
}'
*/
func Create(c *gin.Context) {
	report := &model.Report{}
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	db := container.GetLocalDB()

	report.Status = "init"
	report.CreateTime = time.Now()

	db.Create(&report)

	cases := make([]*model.Case, 0)
	if err := json.Unmarshal(util.InitializeCasesByID("default"), &cases); err != nil {
		fmt.Println(err)
		return
	}

	for idx := range cases {
		cases[idx].Status = "init"

		cases[idx].ReportID = report.ID
		cases[idx].CreateTime = time.Now()
	}

	db.Create(&cases)

	c.JSON(http.StatusOK, gin.H{
		"id":      report.ID,
		"message": "ok",
	})
	return
}
