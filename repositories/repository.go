package repositories

import (
	"../utils"
	mgo "gopkg.in/mgo.v2"
)

func SaveEvent(collection *mgo.Collection, i interface{}) {
	err := collection.Insert(&i)
	utils.CheckAndHandleError(err)
}
