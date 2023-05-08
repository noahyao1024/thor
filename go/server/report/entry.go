package report

import (
	"encoding/json"
	"fmt"
	"golib/server/container"
	"golib/server/model"
	"golib/server/util"
	"time"

	"github.com/gin-gonic/gin"
)

// List ...
func List(c *gin.Context) {
	container.GetLocalDB()

}

// Create ...
func Create(c *gin.Context) {
	report := &model.Report{}
	if err := c.ShouldBindJSON(&report); err != nil {
		fmt.Println(err)
		return
	}

	db := container.GetLocalDB()

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

	fmt.Println(db)
}
