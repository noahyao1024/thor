package cases

import (
	"context"
	"encoding/json"
	"fmt"
	"golib/server/container"
	"golib/server/model"
	"golib/server/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type caseDescription struct {
	Host string `json:"host,omitempty"`
	URI  string `json:"uri,omitempty"`
}

func init() {
}

// BatchRun ...
func BatchRun(c *gin.Context) {
	db := container.GetLocalDB()

	reportID := c.Param("report_id")

	cases := make([]*model.Case, 0)
	db.Model(&model.Case{}).Where("report_id = ? AND status != ?", reportID, "success").Find(&cases)

	var (
		err        error
		succeedCnt int
		failedAt   int
	)

	for _, ca := range cases {
		err = runCase(c, ca)
		if err != nil {
			db.Model(&model.Case{}).Where("id = ?", ca.ID).Update("status", "failure")
			failedAt = ca.ID
			break
		}

		// Update running result
		db.Model(&model.Case{}).Where("id = ?", ca.ID).Update("status", "success")
		succeedCnt++
	}

	if failedAt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":   "failure",
			"failed_at": failedAt,
			"error":     err.Error(),
		})
		return
	}

	db.Model(&model.Report{}).Where("id = ?", reportID).Update("status", "success")

	c.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"succeed_cnt": succeedCnt,
	})
}

func runCase(c *gin.Context, ca *model.Case) error {
	apis := make([]map[string]string, 0)
	json.Unmarshal([]byte(ca.APIs), &apis)

	for _, api := range apis {
		host, _ := api["host"]
		uri, _ := api["uri"]

		assertionType, _ := api["assertion_type"]
		assertionData, _ := api["assertion_data"]

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		ao := &util.APIOption{
			Ctx:  ctx,
			Host: host,
			URI:  uri,
		}

		var data string
		rawResponse, err := util.Call(ao)
		if err != nil {
			return err
		} else {
			dataField := "data"
			data = gjson.GetBytes(rawResponse.Body(), dataField).String()
		}

		switch assertionType {
		default:
			if data != assertionData {
				return fmt.Errorf("expect %s actual %s all %s", assertionData, data, rawResponse.Body())
			}
		}
	}

	return nil
}
