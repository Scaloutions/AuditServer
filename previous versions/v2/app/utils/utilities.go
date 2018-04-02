package utils

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

const (

	// keys for getting actual values
	ACTIVE_ENVIRONMENT = "environment.active"
	DEV                = "development"
	PROD               = "production"
	DOCKER             = "docker"
	DB_HOST            = "database.host"
	DB_PORT            = "database.port"
	EVENT_MAP          = "event-map"
	ENVIRONMENT_MAP    = "env-map"
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

	if v.GetInt(ACTIVE_ENVIRONMENT) == 1 {
		return true
	}
	return false

}

func GetCurrentEnvType(v *viper.Viper) int {
	return v.GetInt(ACTIVE_ENVIRONMENT)
}

func GetEnvPrefix(
	envMap []map[string]interface{}, v *viper.Viper) string {

	curEnvType := GetCurrentEnvType(v)

	switch curEnvType {
	case ToIntFromInt64Inteface(envMap[0]["index"]):
		return envMap[0]["type"].(string)
	case ToIntFromInt64Inteface(envMap[1]["index"]):
		return envMap[1]["type"].(string)
	case ToIntFromInt64Inteface(envMap[2]["index"]):
		return envMap[2]["type"].(string)
	}
	return ""
}

func GetDBURL(v *viper.Viper) string {

	envMap := GetEnvMap(v)

	prefix := GetEnvPrefix(envMap, v)

	var keyArr []interface{}
	keyArr = append(keyArr, prefix)
	keyArr = append(keyArr, ".")
	keyArr = append(keyArr, DB_HOST)
	hostKey := ConcatString(keyArr)
	host := v.GetString(hostKey)
	keyArr = nil
	keyArr = append(keyArr, prefix)
	keyArr = append(keyArr, ".")
	keyArr = append(keyArr, DB_PORT)
	portKey := ConcatString(keyArr)
	port := v.GetString(portKey)
	var urlArr []interface{}
	urlArr = append(urlArr, host)
	urlArr = append(urlArr, ":")
	urlArr = append(urlArr, port)

	return ConcatString(urlArr)
}

func GetEventMap(v *viper.Viper) []map[string]interface{} {
	eventMapI := v.Get(EVENT_MAP)
	return toArrayMap(eventMapI, v)
}

func GetEnvMap(v *viper.Viper) []map[string]interface{} {
	envMapI := v.Get(ENVIRONMENT_MAP)
	return toArrayMap(envMapI, v)
}

func toArrayMap(i interface{}, v *viper.Viper) []map[string]interface{} {

	eventArr, err := toArray(i, v)
	CheckAndHandleError(err)

	eventMapArr := make([]map[string]interface{}, len(eventArr))

	for i := 0; i < len(eventArr); i++ {
		eventMapArr[i], err = toMap(eventArr[i], v)
		CheckAndHandleError(err)
	}

	return eventMapArr

}

func toArray(i interface{}, v *viper.Viper) ([]interface{}, error) {
	arr, ok := i.([]interface{})
	if !ok {
		return nil, errors.New(v.GetString("errors.fail-to-cast-array"))
	}
	return arr, nil
}

func toMap(i interface{}, v *viper.Viper) (map[string]interface{}, error) {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil, errors.New(v.GetString("errors.fail-to-cast-map"))
	}
	return m, nil
}

func ConcatString(list []interface{}) string {
	var buffer bytes.Buffer
	for i := 0; i < len(list); i++ {
		str := fmt.Sprintf("%s", list[i])
		buffer.WriteString(str)
	}
	return buffer.String()
}

func ToIntFromInt64Inteface(i interface{}) int {

	intString := fmt.Sprintf("%d", i)
	integer, err := strconv.ParseInt(intString, 10, 0)
	CheckAndHandleError(err)

	return int(integer)

}
