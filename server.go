package main

/**
reference:
https://medium.com/@maumribeiro/a-fullstack-epic-part-i-a-rest-api
-in-go-accessing-mongo-db-608b46e969cd
*/

import (
	"fmt"

	"./app/controllers"
	"./app/utils"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

const (

	// keys for getting urls
	API                 = "urls.api"
	SYSTEM_EVENT        = "urls.system-event"
	USER_COMMAND        = "urls.user-command"
	QUOTE_SERVER        = "urls.quote-server"
	ERROR_EVENT         = "urls.error-event"
	ACCOUNT_TRANSACTION = "urls.account-transaction"
	LOG_ALL             = "urls.log-all"
	LOG_BY_USER         = "urls.log"

	// other keys
	SERVER_PORT = "environment.port"
)

func getMainEngine(v *viper.Viper) *gin.Engine {

	session := utils.GetDBSession(v)

	eventMap := utils.GetEventMap(v)

	controller :=
		controllers.NewController(session, v, eventMap)

	router := gin.Default()

	api := router.Group(v.GetString(API))
	{
		api.POST(v.GetString(SYSTEM_EVENT), controller.Systemevent)
		api.POST(v.GetString(USER_COMMAND), controller.Usercommand)
		api.POST(v.GetString(QUOTE_SERVER), controller.Quoteserver)
		api.POST(v.GetString(ERROR_EVENT), controller.Errorevent)
		api.POST(v.GetString(ACCOUNT_TRANSACTION), controller.Accounttransaction)
		api.GET(v.GetString(LOG_ALL), controller.LogAll)
		api.GET(v.GetString(LOG_BY_USER), controller.LogByUserName)
		api.GET(v.GetString("urls.clear-db"), controller.ClearDatabase)
	}
	utils.INFO.Println(router)

	return router

}

func main() {

	utils.Init() // initialize loggers
	var viperObj *viper.Viper = utils.SetUpExternalConfig()
	port := viperObj.Get(SERVER_PORT)
	portString := fmt.Sprintf("%d", port)
	router := getMainEngine(viperObj)
	router.Run(":" + portString)

}
