package util

import (
	"context"
	"fmt"
	"log"

	gohttpclient "github.com/bozd4g/go-http-client"
)

// APIOption ...
type APIOption struct {
	Ctx  context.Context
	Host string
	URI  string
}

// Call ...
func Call(a *APIOption) {
	client := gohttpclient.New(a.Host)

	response, err := client.Get(a.Ctx, a.URI)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(response)
}
