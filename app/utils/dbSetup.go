package utils

import (
	"fmt"

	"github.com/spf13/viper"

	mgo "gopkg.in/mgo.v2"
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

// GetDBSession (to be continued)
func GetDBSession(v *viper.Viper) *mgo.Session {
	url := GetDBURL(v)
	session, err := mgo.Dial(url)
	if err != nil {
		errMsg :=
			fmt.Sprintf("Failed to establish connection to Mongo server: \n %s", err)
		ERROR.Println(errMsg)
	}
	INFO.Println(SUCCESSFULCONNECTTODBMSG)
	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)
	return session
}

// GetEventCollection (to be continued)
func GetEventCollection(session *mgo.Session) *mgo.Collection {
	collection := session.DB(DBNAME).C(DOCNAME)
	return collection
}
