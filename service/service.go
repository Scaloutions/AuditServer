package service

import (
	"encoding/json"
	"io/ioutil"

	"../data"
	"../utils"
	"github.com/gin-gonic/gin"
)

func checkAndHandleError(err error) {
	if err != nil {
		utils.ERROR.Println(err)
	}
}

func getBody(c *gin.Context) []byte {
	body, err := ioutil.ReadAll(c.Request.Body)
	checkAndHandleError(err)
	return body
}

func getObject(i interface{}, body []byte) interface{} {
	err := json.Unmarshal(body, &i)
	checkAndHandleError(err)
	return i
}

func logXMLObj(i interface{}) {
	xmlEvent := utils.GetXMLEventString(i)
	utils.XMLLOGGER.Println(xmlEvent)
}

func processingHelper(body []byte, i interface{}) {
	getObject(&i, body)
	logXMLObj(i) // instead of logging, sending data to database
}

func Processing(
	commandType string,
	c *gin.Context) {

	body := getBody(c)
	utils.INFO.Println(string(body))

	switch commandType {
	case "usercommand":
		var userCommand data.UserCommand
		processingHelper(body, &userCommand)
	case "systemevent":
		var systemEvent data.SystemEvent
		processingHelper(body, &systemEvent)
	case "accounttransaction":
		var accountTransaction data.AccountTransaction
		processingHelper(body, &accountTransaction)
	case "quoteserver":
		var quoteServer data.QuoteServer
		processingHelper(body, &quoteServer)
	case "errorevent":
		var errorEvent data.ErrorEvent
		processingHelper(body, &errorEvent)
	}
}
