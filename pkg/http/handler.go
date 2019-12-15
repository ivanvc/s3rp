package http

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ivanvc/s3-proxy/pkg/config"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/ivanvc/s3-proxy/pkg/log"
)

type handler struct {
	s3svc *s3.S3
}

func newHandler(sess *session.Session) *handler {
	svc := s3.New(sess)
	return &handler{svc}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	status := http.StatusOK
	resp, err := h.s3svc.GetObject(&s3.GetObjectInput{
		Bucket: &config.Bucket,
		Key:    h.getKey(req.URL),
	})
	defer func() {
		if err != nil {
			log.Logger.Printf("%d %v %s %s [%v]\n", status, time.Now().Sub(t), req.Method, req.URL.Path, err)
		} else {
			log.Logger.Printf("%d %v %s %s\n", status, time.Now().Sub(t), req.Method, req.URL.Path)
		}
	}()

	if err != nil {
		status = http.StatusBadRequest
		http.Error(w, err.Error(), status)
		return
	}

	defer resp.Body.Close()
	w.Header().Set("Content-Type", *resp.ContentType)
	io.Copy(w, resp.Body)
}

func (h *handler) getKey(url *url.URL) *string {
	key := url.Path
	if strings.HasSuffix(key, "/") {
		key += config.IndexPage
	}
	key = key[1:len(key)]
	return &key
}
