package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"./utils"
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
	logXMLObj(i)
}

func processing(
	commandType string,
	c *gin.Context) {

	body := getBody(c)
	utils.INFO.Println(string(body))

	switch commandType {
	case "usercommand":
		var userCommand utils.UserCommand
		processingHelper(body, &userCommand)
	case "systemevent":
		var systemEvent utils.SystemEvent
		processingHelper(body, &systemEvent)
	case "accounttransaction":
		var accountTransaction utils.AccountTransaction
		processingHelper(body, &accountTransaction)
	case "quoteserver":
		var quoteServer utils.QuoteServer
		processingHelper(body, &quoteServer)
	case "errorevent":
		var errorEvent utils.ErrorEvent
		processingHelper(body, &errorEvent)
	}

}

func getMainEngine() *gin.Engine {

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/systemevent", systemevent)
		api.POST("/usercommand", usercommand)
		api.POST("/quoteserver", quoteserver)
		api.POST("/errorevent", errorevent)
		api.POST("/accounttransaction", accounttransaction)
	}
	return router

}

func main() {

	utils.Init() // initialize loggers

	router := getMainEngine()

	router.Run(":8082")
}
