package http

import (
	"net/url"
	"testing"

	"github.com/ivanvc/s3-reverse-proxy/pkg/config"
	. "github.com/ivanvc/s3-reverse-proxy/pkg/internal/test"
	"github.com/ivanvc/s3-reverse-proxy/pkg/log"

	"github.com/aws/aws-sdk-go/aws/session"
)

func init() {
	log.Logger = EmptyLogger
}

func TestGetURL(t *testing.T) {
	config.Bucket = "my-bucket"
	config.IndexPage = "index.html"
	p := newProxy(new(session.Session))
	u, _ := url.Parse("http://localhost/")

	result := p.getURL(u)
	if result.Path != "/index.html" {
		t.Errorf("IndexPage not appended (%s)", result)
	}
	if result.Host != "my-bucket.s3.amazonaws.com" {
		t.Errorf("Bucket not set (%s)", result)
	}
	u.Path = "/test"
	if result = p.getURL(u); result.Path != "/test" {
		t.Errorf("Path modified (%s)", result)
	}
	config.IndexPage = "test.php3"
	u.Path = "/test/"
	if result = p.getURL(u); result.Path != "/test/test.php3" {
		t.Errorf("IndexPage not appended (%s)", result)
	}
}
