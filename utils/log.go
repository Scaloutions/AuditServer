package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	TRACE   *log.Logger
	INFO    *log.Logger
	WARNING *log.Logger
	ERROR   *log.Logger
)

func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	TRACE = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	INFO = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	WARNING = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ERROR = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Init() {
	logfile, err := os.Create("goag.log")
	if err != nil {
		log.Fatalln("Fail to open log file")
	}
	initLog(ioutil.Discard, logfile, logfile, logfile)
}
