package utils

import (
	"log"
)

var LogWriter FakeWriter = NewFakeWriter()
var Logger *log.Logger = log.New(LogWriter, "utils.Writer: ", log.LstdFlags)