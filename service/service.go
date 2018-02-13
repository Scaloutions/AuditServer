package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/mgo.v2/bson"

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

	eventMap := data.GetEventMap()

	switch i.(type) {
	case *data.UserCommandEvent:
		userCommandEvent := i.(*data.UserCommandEvent)
		userCommandEvent.ID = bson.NewObjectId()
		userCommandEvent.EventType = eventMap["usercommad"]
		userCommandEvent.Timestamp = utils.GetCurrentTs()
		repositories.SaveEvent(collection, userCommandEvent)
	case *data.SystemEventJ:
		systemEventJ := i.(*data.SystemEventJ)
		systemEventJ.ID = bson.NewObjectId()
		systemEventJ.EventType = eventMap["systemevent"]
		systemEventJ.Timestamp = utils.GetCurrentTs()
		repositories.SaveEvent(collection, systemEventJ)
	case *data.AccountTransactionEvent:
		accountTransactionEvent := i.(*data.AccountTransactionEvent)
		accountTransactionEvent.ID = bson.NewObjectId()
		accountTransactionEvent.EventType = eventMap["accounttransaction"]
		accountTransactionEvent.Timestamp = utils.GetCurrentTs()
		repositories.SaveEvent(collection, accountTransactionEvent)
	case *data.QuoteServerEvent:
		quoteServerEvent := i.(*data.QuoteServerEvent)
		quoteServerEvent.ID = bson.NewObjectId()
		quoteServerEvent.EventType = eventMap["quoteserver"]
		quoteServerEvent.Timestamp = utils.GetCurrentTs()
		repositories.SaveEvent(collection, quoteServerEvent)
	case *data.ErrorEventJ:
		errorEventJ := i.(*data.ErrorEventJ)
		errorEventJ.ID = bson.NewObjectId()
		errorEventJ.EventType = eventMap["errorevent"]
		errorEventJ.Timestamp = utils.GetCurrentTs()
		repositories.SaveEvent(collection, errorEventJ)
	}
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

func LogAll(colllection *mgo.Collection) {

	eventMap := data.GetEventMap()
	var results []map[string]interface{}
	error := colllection.Find(nil).All(&results)
	utils.CheckAndHandleError(error)
	for _, event := range results {
		processXMLEvent(event, eventMap)
	}
}

func LogByUserName(collection *mgo.Collection, c *gin.Context) {

	eventMap := data.GetEventMap()
	userName := c.Param("userName")
	fmt.Println(userName)
	var results []map[string]interface{}
	error := collection.Find(bson.M{"username": userName}).All(&results)
	utils.CheckAndHandleError(error)
	for _, event := range results {
		processXMLEvent(event, eventMap)
	}
}

func processXMLEvent(event map[string]interface{}, eventMap map[string]int) {

	eventType, _ := event["eventtype"].(int)
	server, _ := event["server"].(string)
	transactionNum, _ := event["transactionnum"].(int)
	usrName, _ := event["username"].(string)
	timestamp, _ := event["timestamp"].(int64)

	switch eventType {
	case eventMap["usercommand"]:
		command, _ := event["command"].(string)
		stockSymbol, _ := event["stocksymbol"].(string)
		funds, _ := event["funds"].(float64)
		userCommand := data.GetUserCommand(
			server,
			transactionNum,
			command,
			usrName,
			stockSymbol,
			funds,
			timestamp)
		logXMLObj(userCommand)
	case eventMap["systemevent"]:
		command, _ := event["command"].(string)
		stockSymbol, _ := event["stocksymbol"].(string)
		funds, _ := event["funds"].(float64)
		systemEvent := data.GetSystemEvent(
			server,
			transactionNum,
			command,
			usrName,
			stockSymbol,
			funds,
			timestamp)
		logXMLObj(systemEvent)
	case eventMap["errorevent"]:
		command, _ := event["command"].(string)
		stockSymbol, _ := event["stocksymbol"].(string)
		funds, _ := event["funds"].(float64)
		errMsg, _ := event["errormessage"].(string)
		errorEvent := data.GetErrorEvent(
			server,
			transactionNum,
			command,
			usrName,
			stockSymbol,
			funds,
			errMsg,
			timestamp)
		logXMLObj(errorEvent)
	case eventMap["quoteserver"]:
		quoteServerTime, _ := event["quoteservertime"].(int64)
		command, _ := event["command"].(string)
		stockSymbol, _ := event["stocksymbol"].(string)
		price, _ := event["price"].(float64)
		cryptokey, _ := event["cryptokey"].(string)
		quoteServer := data.GetQuoteServer(
			server,
			transactionNum,
			quoteServerTime,
			command,
			usrName,
			stockSymbol,
			price,
			cryptokey,
			timestamp)
		logXMLObj(quoteServer)
	case eventMap["accounttransaction"]:
		action, _ := event["action"].(string)
		funds, _ := event["funds"].(float64)
		accountTransaction := data.GetAccountTransaction(
			server,
			transactionNum,
			action,
			usrName,
			funds,
			timestamp)
		logXMLObj(accountTransaction)
	}
}
