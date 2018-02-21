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

func getMainEngine(v *viper.Viper) *gin.Engine {

	session := utils.GetDBSession(v)
	log.Println(session)

	controller :=
		controllers.NewController(session, v)

	router := gin.Default()

	api := router.Group(v.GetString("urls.api"))
	{
		api.POST(v.GetString("urls.system-event"), controller.Systemevent)
		api.POST(v.GetString("urls.user-command"), controller.Usercommand)
		api.POST(v.GetString("urls.quote-server"), controller.Quoteserver)
		api.POST(v.GetString("urls.error-event"), controller.Errorevent)
		api.POST(v.GetString("urls.account-transaction"), controller.Accounttransaction)
		api.POST(v.GetString("urls.log-all"), controller.LogAll)
		api.POST(v.GetString("urls.log"), controller.LogByUserName)
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
