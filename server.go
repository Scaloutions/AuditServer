package main

import (
	"./utils"
)

func main() {

	utils.Init() // initialize loggers

	session := utils.GetDBSession()
	defer session.Close()

	router := utils.GetMainEngine()

	router.Run(":8082")
}
