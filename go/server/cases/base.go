package cases

import (
	"context"
	"encoding/json"
	"fmt"
	"golib/server/container"
	"golib/server/model"
	"golib/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type caseDescription struct {
	Host string `json:"host,omitempty"`
	URI  string `json:"uri,omitempty"`
}

// Cases ...
var Cases []caseDescription

func init() {
	Cases = make([]caseDescription, 0)
	raw := `
	[
		{
			"host": "https://jsonplaceholder.typicode.com",
			"uri": "/posts/1"
		}
	]
	`
	json.Unmarshal([]byte(raw), &Cases)
}

// BatchRun ...
func BatchRun(c *gin.Context) {
	db := container.GetLocalDB()

	reportID := c.Param("report_id")

	cases := make([]*model.Case, 0)
	db.Where(&model.Case{}).Where("report_id = ?", reportID).Find(&cases)

	for _, ca := range cases {
		err := runCase(c, ca)
		if err != nil {

		}

		// TODO Check assertion

		// TODO Update running result
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
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

		ao := &util.APIOption{
			Ctx:  context.Background(),
			Host: host,
			URI:  uri,
		}

		var data string

		rawResponse, err := util.Call(ao)
		if err == nil {
			dataField := "data"
			data = gjson.GetBytes(rawResponse.Body(), dataField).String()
		} else {
			return err
		}

		switch assertionType {
		default:
			if data != assertionData {
				return fmt.Errorf("expect %s actual %s all %s", assertionData, data, rawResponse.Body())
			}
			panic("ss")
		}
	}

	return nil
}
