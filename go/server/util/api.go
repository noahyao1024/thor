package util

import (
	"context"
	"time"

	gohttpclient "github.com/bozd4g/go-http-client"
)

// APIOption ...
type APIOption struct {
	Ctx  context.Context
	Host string
	URI  string
}

// Call ...
func Call(a *APIOption) (*gohttpclient.Response, error) {
	client := gohttpclient.New(a.Host, gohttpclient.WithTimeout(1*time.Second))

	response, err := client.Get(a.Ctx, a.URI)
	if err != nil {
		return response, err
	}

	return response, nil
}
