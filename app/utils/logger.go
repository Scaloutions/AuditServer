package utils

import (
	"io"
	"log"
	"os"
)

const (
	TRACE_FILE_PATH   = "trace.log"
	INFO_FILE_PATH    = "info.log"
	WARNING_FILE_PATH = "warning.log"
	ERROR_FILE_PATH   = "error.log"
	LOG_FILE_PATH     = "1UserLogFile.xml"
)

var (
	TRACE     *log.Logger
	INFO      *log.Logger
	WARNING   *log.Logger
	ERROR     *log.Logger
	XMLLOGGER *log.Logger
)

func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	xmlHandle io.Writer) {

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

	XMLLOGGER = log.New(xmlHandle, "", 0)
}

// Init To initialize loggers
func Init() {
	traceLogFile, err1 := os.Create(TRACE_FILE_PATH)
	infoLogFile, err2 := os.Create(INFO_FILE_PATH)
	warningLogFile, err3 := os.Create(WARNING_FILE_PATH)
	errorLogFile, err4 := os.Create(ERROR_FILE_PATH)
	xmlLogFile, err5 := os.Create(LOG_FILE_PATH)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		ERROR.Fatalln("Fail to open log file")
	}
	initLog(traceLogFile, infoLogFile, warningLogFile, errorLogFile, xmlLogFile)
	XMLLOGGER.Println("<?xml version=\"1.0\"?>")
}
