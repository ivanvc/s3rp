package http

import (
	"net/http"

	"github.com/ivanvc/s3-reverse-proxy/pkg/config"
	"github.com/ivanvc/s3-reverse-proxy/pkg/log"

	"github.com/aws/aws-sdk-go/aws/session"
)

// Serve starts the HTTP server.
func Serve(sess *session.Session) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/", newProxy(sess))
	log.Logger.Println("Ready, listening on", config.Host)
	log.Logger.Fatal(http.ListenAndServe(config.Host, mux))
}
