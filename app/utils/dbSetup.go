package utils

import (
	"fmt"

	"github.com/spf13/viper"

	mgo "gopkg.in/mgo.v2"
)

const (

	// DBNAME the name of the DB instance
	DBNAME = "eventstore"
	// DOCNAME the name of the document
	DOCNAME = "events"
)

// GetDBSession (to be continued)
func GetDBSession(v *viper.Viper) *mgo.Session {
	url := GetDBURL(v)

	session, err := mgo.Dial(url)
	if err != nil {
		errMsg :=
			fmt.Sprintf(v.GetString("errors.fail-to-connect-db-server"), err)
		ERROR.Println(errMsg)
	}
	INFO.Println(v.GetString("database:successful-connection-to-db"))
	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)
	return session
}

// GetEventCollection (to be continued)
func GetEventCollection(session *mgo.Session) *mgo.Collection {
	collection := session.DB(DBNAME).C(DOCNAME)
	return collection
}
