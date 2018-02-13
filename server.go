package main

/**
reference:
https://medium.com/@maumribeiro/a-fullstack-epic-part-i-a-rest-api
-in-go-accessing-mongo-db-608b46e969cd
*/

import (
	"./controller"
	"./utils"
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
)

func getMainEngine() *gin.Engine {

	router := gin.Default()

	api := router.Group(APIURL)
	{
		api.POST(SYSEVENTURL, controller.Systemevent)
		api.POST(USRCOMMANDURL, controller.Usercommand)
		api.POST(QUOSERVERURL, controller.Quoteserver)
		api.POST(ERREVENTURL, controller.Errorevent)
		api.POST(ACCTTRANSACTIONURL, controller.Accounttransaction)
	}
	return router

}

func main() {

	utils.Init() // initialize loggers

	router := getMainEngine()

	router.Run(":8082")
}
