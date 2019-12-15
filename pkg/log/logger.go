package log

import (
	"io/ioutil"
	"log"
)

var Logger *log.Logger

func init() {
	Logger = log.New(ioutil.Discard, "", 0)
}
