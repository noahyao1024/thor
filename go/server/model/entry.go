package model

import "time"

// Report ...
type Report struct {
	ID         int       `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Data       string    `json:"data,omitempty"`
	Status     string    `json:"status,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	ModifyTime time.Time `json:"modify_time,omitempty"`
}

// TableName ...
func (i *Report) TableName() string {
	return "reports"
}

// ReportData ...
type ReportData struct {
	Version string `json:"version,omitempty"`
}

// Case ...
type Case struct {
	ID           int       `json:"id,omitempty"`
	TemplateID   int       `json:"template_id,omitempty"`
	TemplateName string    `json:"template_name,omitempty"`
	ReportID     int       `json:"report_id,omitempty"`
	Status       string    `json:"status,omitempty"`
	Data         string    `json:"data,omitempty"`
	APIs         string    `json:"apis,omitempty"`
	CreateTime   time.Time `json:"create_time,omitempty"`
	ModifyTime   time.Time `json:"modify_time,omitempty"`
}

// TableName ...
func (i *Case) TableName() string {
	return "cases"
}
