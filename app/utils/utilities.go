package utils

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/spf13/viper"
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

func CheckAndHandleError(err error) {
	if err != nil {
		ERROR.Println(err)
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

func IsDevEnv(v *viper.Viper) bool {
	if v.GetInt("environment.active") == 1 {
		return true
	}
	return false

}

func GetDBURL(v *viper.Viper) string {

	var prefix string
	if IsDevEnv(v) {
		prefix = "development"
	} else {
		prefix = "production"
	}
	var keyArr []interface{}
	keyArr = append(keyArr, prefix)
	keyArr = append(keyArr, ".")
	keyArr = append(keyArr, "database.host")
	hostKey := ConcatString(keyArr)
	host := v.GetString(hostKey)
	keyArr = nil
	keyArr = append(keyArr, prefix)
	keyArr = append(keyArr, ".")
	keyArr = append(keyArr, "database.port")
	portKey := ConcatString(keyArr)
	port := v.GetString(portKey)
	var urlArr []interface{}
	urlArr = append(urlArr, host)
	urlArr = append(urlArr, ":")
	urlArr = append(urlArr, port)

	return ConcatString(urlArr)
}

func ConcatString(list []interface{}) string {

	var buffer bytes.Buffer

	for i := 0; i < len(list); i++ {
		str := fmt.Sprintf("%s", list[i])
		buffer.WriteString(str)
	}

	return buffer.String()

}
