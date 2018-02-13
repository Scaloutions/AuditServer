package main

/**
reference:
https://medium.com/@maumribeiro/a-fullstack-epic-part-i-a-rest-api
-in-go-accessing-mongo-db-608b46e969cd
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"./data"
	"./utils"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

const (

	/*
		DB related
	*/

	// SERVER the DB server
	SERVER = "localhost:27017"
	// DBNAME the name of the DB instance
	DBNAME = "eventstore"
	// DOCNAME the name of the document
	DOCNAME = "events"

	/*
		URLs
	*/

	// APIURL (to be continued)
	APIURL = "/api"
	// SYSEVENTURL (to be continued)
	SYSEVENTURL = "/systemevent"
	// USRCOMMANDURL (to be continued)
	USRCOMMANDURL = "/usercommand"
	// QUOSERVERURL (to be continued)
	QUOSERVERURL = "/quoteserver"
	// ERREVENTURL (to be continued)
	ERREVENTURL = "/errorevent"
	// ACCTTRANSACTIONURL (to be continued)
	ACCTTRANSACTIONURL = "/accounttransaction"

	/*
		Messages
	*/

	// SUCCESSFULCONNECTTODBMSG (to be continued)
	SUCCESSFULCONNECTTODBMSG = "Connecting to mongodb server successfully..."
)

func getDBSession() *mgo.Session {

	session, err := mgo.Dial(SERVER)
	if err != nil {
		errMsg :=
			fmt.Sprintf("Failed to establish connection to Mongo server: \n %s", err)
		utils.ERROR.Println(errMsg)
	}
	utils.INFO.Println(SUCCESSFULCONNECTTODBMSG)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

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

func usercommand(c *gin.Context) {
	processing("usercommand", c)
}

func systemevent(c *gin.Context) {
	processing("systemevent", c)
}

func quoteserver(c *gin.Context) {
	processing("quoteserver", c)
}

func accounttransaction(c *gin.Context) {
	processing("accounttransaction", c)
}

func errorevent(c *gin.Context) {
	processing("errorevent", c)
}

func processingHelper(body []byte, i interface{}) {
	getObject(&i, body)
	//logXMLObj(i) // instead of logging, sending data to database
}

func processing(
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

func getMainEngine() *gin.Engine {

	router := gin.Default()

	api := router.Group(APIURL)
	{
		api.POST(SYSEVENTURL, systemevent)
		api.POST(USRCOMMANDURL, usercommand)
		api.POST(QUOSERVERURL, quoteserver)
		api.POST(ERREVENTURL, errorevent)
		api.POST(ACCTTRANSACTIONURL, accounttransaction)
	}
	return router

}

func main() {

	utils.Init() // initialize loggers

	session := getDBSession()
	defer session.Close()

	router := getMainEngine()

	router.Run(":8082")
}
