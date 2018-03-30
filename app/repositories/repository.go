package repositories

import (
	"log"

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
	err := collection.Find(nil).All(&results)
	utils.CheckAndHandleError(err)

	return results
}

func GetAllEventsByUser(
	collection *mgo.Collection,
	userName string) []map[string]interface{} {

	var results []map[string]interface{}
	err := collection.Find(bson.M{"username": userName}).All(&results)
	utils.CheckAndHandleError(err)
	return results

}

func RemoveAll(collection *mgo.Collection) {

	info, err := collection.RemoveAll(nil)
	utils.CheckAndHandleError(err)
	log.Println(info)
}
