package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

/**
reference:
https://medium.com/@maumribeiro/a-fullstack-epic-part-i-a-rest-api
-in-go-accessing-mongo-db-608b46e969cd
*/

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

// GetDBSession Set up database connection & return session
func GetDBSession() *mgo.Session {

	session, err := mgo.Dial(SERVER)
	if err != nil {
		errMsg :=
			fmt.Sprintf("Failed to establish connection to Mongo server: \n %s", err)
		ERROR.Println(errMsg)
	}
	INFO.Println(SUCCESSFULCONNECTTODBMSG)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func checkAndHandleError(err error) {
	if err != nil {
		ERROR.Println(err)
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
	xmlEvent := GetXMLEventString(i)
	XMLLOGGER.Println(xmlEvent)
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
	INFO.Println(string(body))

	switch commandType {
	case "usercommand":
		var userCommand UserCommand
		processingHelper(body, &userCommand)
	case "systemevent":
		var systemEvent SystemEvent
		processingHelper(body, &systemEvent)
	case "accounttransaction":
		var accountTransaction AccountTransaction
		processingHelper(body, &accountTransaction)
	case "quoteserver":
		var quoteServer QuoteServer
		processingHelper(body, &quoteServer)
	case "errorevent":
		var errorEvent ErrorEvent
		processingHelper(body, &errorEvent)
	}

}

// GetMainEngine (to be continued)
func GetMainEngine() *gin.Engine {

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
