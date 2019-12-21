package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/ivanvc/s3-reverse-proxy/pkg/config"
	"github.com/ivanvc/s3-reverse-proxy/pkg/log"

	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

type proxy struct {
	httputil.ReverseProxy
	session    *session.Session
	clientInfo *metadata.ClientInfo
}

func newProxy(sess *session.Session) *proxy {
	p := &proxy{
		session:    sess,
		clientInfo: &metadata.ClientInfo{ServiceName: "s3"},
	}
	p.Director = p.getDirector
	return p
}

func (p *proxy) getDirector(req *http.Request) {
	defer log.Logger.Printf("%s %s\n", req.Method, req.URL.Path)

	req.URL = p.getURL(req.URL)
	req.Host = req.URL.Host

	headers, err := p.signedHeaders(req)
	if err != nil {
		log.Logger.Println("Error signing request", err)
	}
	for k, v := range headers {
		req.Header[k] = v
	}
}

func (p *proxy) signedHeaders(req *http.Request) (map[string][]string, error) {
	operation := &request.Operation{
		Name:       "GetObject",
		HTTPPath:   req.URL.Path,
		HTTPMethod: http.MethodGet,
	}

	handlers := request.Handlers{}
	handlers.Sign.PushBackNamed(v4.SignRequestHandler)

	awsReq := request.New(*p.session.Config, *p.clientInfo, handlers, nil, operation, nil, nil)
	awsReq.HTTPRequest.URL = req.URL
	if err := awsReq.Sign(); err != nil {
		return nil, err
	}
	return awsReq.HTTPRequest.Header, nil
}

func (proxy) getURL(url *url.URL) *url.URL {
	key := url.Path
	if strings.HasSuffix(key, "/") {
		key += config.IndexPage
	}
	key = key[1:len(key)]
	u, err := url.Parse("http://" + config.Bucket + ".s3.amazonaws.com/" + key)
	if err != nil {
		log.Logger.Println("Error generating URL", err.Error())
	}
	return u
}
