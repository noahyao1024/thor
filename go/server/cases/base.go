package cases

import (
	"context"
	"encoding/json"
	"fmt"
	"golib/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
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

	fmt.Println(Cases)
}

// Wrap ...
func Wrap(c *gin.Context) {
	for _, ca := range Cases {
		ao := &util.APIOption{
			Ctx:  context.Background(),
			Host: ca.Host,
			URI:  ca.URI,
		}

		var (
			response interface{}
			message  string
		)

		rawResponse, err := util.Call(ao)
		if err == nil {
			response = string(rawResponse.Body())
		} else {
			message = err.Error()
		}

		c.JSON(http.StatusOK, gin.H{
			"response": response,
			"message":  message,
		})

		// TODO Check assertion

		// TODO Update running result
	}
}
