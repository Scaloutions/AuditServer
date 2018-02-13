package service

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/mgo.v2"

	"../data"
	"../repositories"
	"../utils"
	"github.com/gin-gonic/gin"
)

func getBody(c *gin.Context) []byte {
	body, err := ioutil.ReadAll(c.Request.Body)
	utils.CheckAndHandleError(err)
	return body
}

func getObject(i interface{}, body []byte) interface{} {
	err := json.Unmarshal(body, &i)
	utils.CheckAndHandleError(err)
	return i
}

func logXMLObj(i interface{}) {
	xmlEvent := utils.GetXMLEventString(i)
	utils.XMLLOGGER.Println(xmlEvent)
}

func processingHelper(
	body []byte,
	i interface{},
	collection *mgo.Collection) {

	getObject(&i, body)
	repositories.SaveEvent(collection, i)
}

func Processing(commandType string, c *gin.Context, collection *mgo.Collection) {

	body := getBody(c)
	utils.INFO.Println(string(body))

	switch commandType {
	case "usercommand":
		var userCommandEvent data.UserCommandEvent
		processingHelper(body, &userCommandEvent, collection)
	case "systemevent":
		var systemEventJ data.SystemEventJ
		processingHelper(body, &systemEventJ, collection)
	case "accounttransaction":
		var accountTransactionEvent data.AccountTransactionEvent
		processingHelper(body, &accountTransactionEvent, collection)
	case "quoteserver":
		var quoteServerEvent data.QuoteServerEvent
		processingHelper(body, &quoteServerEvent, collection)
	case "errorevent":
		var errorEventJ data.ErrorEventJ
		processingHelper(body, &errorEventJ, collection)
	}
}
