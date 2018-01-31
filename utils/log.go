package utils

import (
	"io"
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

// Init To initialize loggers
func Init() {
	traceLogFile, err1 := os.Create("trace.log")
	infoLogFile, err2 := os.Create("info.log")
	warningLogFile, err3 := os.Create("warning.log")
	errorLogFile, err4 := os.Create("error.log")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Fatalln("Fail to open log file")
	}
	initLog(traceLogFile, infoLogFile, warningLogFile, errorLogFile)
}
