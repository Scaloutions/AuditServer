package repositories

import (
	"../utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SaveEvent(collection *mgo.Collection, i interface{}) {
	err := collection.Insert(&i)
	utils.CheckAndHandleError(err)
}

func GetAllEvents(
	collection *mgo.Collection) []map[string]interface{} {

	var results []map[string]interface{}
	error := collection.Find(nil).All(&results)
	utils.CheckAndHandleError(error)

	return results
}

func GetAllEventsByUser(
	collection *mgo.Collection,
	userName string) []map[string]interface{} {

	var results []map[string]interface{}
	error := collection.Find(bson.M{"username": userName}).All(&results)
	utils.CheckAndHandleError(error)
	return results

}
