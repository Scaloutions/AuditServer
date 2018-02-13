package repository

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"

	"../utils"
)

const (
	// SERVER the DB server
	SERVER = "localhost:27017"
	// DBNAME the name of the DB instance
	DBNAME = "eventstore"
	// DOCNAME the name of the document
	DOCNAME = "events"

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

// CreateTable (to be continued)
func CreateTable() *mgo.Collection {
	session := getDBSession()
	defer session.Close()
	collection := session.DB(DBNAME).C(DOCNAME)
	return collection
}
