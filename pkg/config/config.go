package config

import (
	"flag"
	"os"

	"github.com/ivanvc/s3-proxy/pkg/log"
)

var (
	indexPage = flag.String("index", "index.html", "index page (default: index.html)")
	bucket    = flag.String("bucket", "", "bucket name")
	host      = flag.String("host", ":8080", "host to listen on (default: \":8080\")")
	IndexPage string
	Bucket    string
	Host      string
)

func init() {
	flag.Parse()
}

func Load() {
	if v := os.Getenv("BUCKET"); len(v) > 0 {
		Bucket = v
	} else {
		Bucket = *bucket
	}
	if v := os.Getenv("HOST"); len(v) > 0 {
		Host = v
	} else {
		Host = *host
	}
	if v := os.Getenv("INDEX_PAGE"); len(v) > 0 {
		IndexPage = v
	} else {
		IndexPage = *indexPage
	}

	if len(Bucket) == 0 {
		log.Logger.Fatal("BUCKET not set")
	}
}
