package util

import (
	"encoding/json"
	"golib/server/model"
	"time"
)

type internalCase struct {
	ID         int           `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	ReportID   int           `json:"report_id,omitempty"`
	Status     string        `json:"status,omitempty"`
	Data       string        `json:"data,omitempty"`
	APIs       []interface{} `json:"apis,omitempty"`
	CreateTime time.Time     `json:"create_time,omitempty"`
	ModifyTime time.Time     `json:"modify_time,omitempty"`
}

// InitializeCasesByID ...
func InitializeCasesByID(name string) []byte {
	raw := `
	[
		{
		  "name": "1-123",
		  "id": 123,
		  "apis": [
			{
			  "host": "http://10.11.1.3",
			  "uri": "/system/alive",
			  "assertion_data": "{\"open\":{},\"sems-vr\":{},\"shepherd\":{}}"
			}
		  ]
		}
	]
	`

	internalCases := make([]*internalCase, 0)
	json.Unmarshal([]byte(raw), &internalCases)

	cases := make([]*model.Case, 0)
	for _, internalCase := range internalCases {
		ca := &model.Case{
			TemplateID:   internalCase.ID,
			TemplateName: internalCase.Name,
			ReportID:     internalCase.ReportID,
			Status:       internalCase.Status,
			Data:         internalCase.Data,

			CreateTime: internalCase.CreateTime,
			ModifyTime: internalCase.ModifyTime,
		}

		bytesOfAPIs, _ := json.Marshal(internalCase.APIs)
		ca.APIs = string(bytesOfAPIs)

		cases = append(cases, ca)
	}

	bytes, _ := json.Marshal(cases)
	return bytes
}
