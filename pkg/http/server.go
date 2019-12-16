package http

import (
	"net/http"

	"github.com/ivanvc/s3-proxy/pkg/config"
	"github.com/ivanvc/s3-proxy/pkg/log"

	"github.com/aws/aws-sdk-go/aws/session"
)

// Serve starts the HTTP server.
func Serve(sess *session.Session) {
	log.Logger.Println("Ready, listening on", config.Host)
	log.Logger.Printf("Listening on %v\n", config.Host)
	log.Logger.Fatal(http.ListenAndServe(config.Host, newProxy(sess)))
}
