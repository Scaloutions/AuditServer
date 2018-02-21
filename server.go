package main

/**
reference:
https://medium.com/@maumribeiro/a-fullstack-epic-part-i-a-rest-api
-in-go-accessing-mongo-db-608b46e969cd
*/

import (
	"log"

	"./app/controllers"
	"./app/utils"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

const (

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
	// LOGALLURL (to be continued)
	LOGALLURL = "/log"
	// LOGURL (to be continued)
	LOGURL = "/log/:userName"
)

func getMainEngine(v *viper.Viper) *gin.Engine {

	session := utils.GetDBSession(v)
	log.Println(session)

	controller :=
		controllers.NewController(session)

	router := gin.Default()

	api := router.Group(APIURL)
	{
		api.POST(SYSEVENTURL, controller.Systemevent)
		api.POST(USRCOMMANDURL, controller.Usercommand)
		api.POST(QUOSERVERURL, controller.Quoteserver)
		api.POST(ERREVENTURL, controller.Errorevent)
		api.POST(ACCTTRANSACTIONURL, controller.Accounttransaction)
		api.POST(LOGALLURL, controller.LogAll)
		api.POST(LOGURL, controller.LogByUserName)
	}
	utils.INFO.Println(router)

	return router

}

func main() {

	utils.Init() // initialize loggers
	viperObj := utils.SetUpExternalConfig()
	router := getMainEngine(viperObj)

	router.Run(":8082")
}
