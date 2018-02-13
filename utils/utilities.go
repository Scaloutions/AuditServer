package utils

import (
	"encoding/xml"
	"fmt"
	"time"
)

func GetCurrentTs() int64 {
	return time.Now().UnixNano() / 1000000
}

func GetFundsAsString(amount float64) string {
	if amount == 0 {
		return ""
	}
	return fmt.Sprintf("%.2f", float64(amount))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetXMLEventString(loggingObject interface{}) string {

	var xmlString string
	if xmlstring, err := xml.MarshalIndent(loggingObject, "", "    "); err == nil {
		xmlString = string(xmlstring)
		return xmlString
	}
	return xmlString

}
