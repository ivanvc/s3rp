package main

import (
	"github.com/ivanvc/s3-proxy/pkg/config"
	"github.com/ivanvc/s3-proxy/pkg/http"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	config.Load()
	http.Serve(session.Must(session.NewSession()))
}
