package utils

import (
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
	} else {
		return false
	}
}

func GetDBURL(v *viper.Viper) string {

	var prefix string
	if IsDevEnv(v) {
		prefix = "development"
	} else {
		prefix = "production"
	}
	key := prefix + "." + "database.url"
	return v.GetString(key)
}
