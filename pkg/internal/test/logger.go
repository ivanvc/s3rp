package test

import (
	"io/ioutil"
	"log"
)

var EmptyLogger *log.Logger = log.New(ioutil.Discard, "", log.Lshortfile)
