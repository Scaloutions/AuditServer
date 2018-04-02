package controllers

import (
	"../service"
	"../utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

type (
	Controller struct {
		session  *mgo.Session
		viperObj *viper.Viper
		eventMap []map[string]interface{}
	}
)

const (
	EVENT_TYPE = "event-type"
)

func NewController(
	s *mgo.Session, v *viper.Viper, eventMap []map[string]interface{}) *Controller {
	return &Controller{s, v, eventMap}
}

func (controller Controller) Usercommand(c *gin.Context) {
	index := 0
	helperFunc(controller, index, c)
}

func (controller Controller) Systemevent(c *gin.Context) {

	index := 1
	helperFunc(controller, index, c)
}

func (controller Controller) Quoteserver(c *gin.Context) {

	index := 3
	helperFunc(controller, index, c)
}

func (controller Controller) Accounttransaction(c *gin.Context) {

	index := 4
	helperFunc(controller, index, c)
}

func (controller Controller) Errorevent(c *gin.Context) {

	index := 2
	helperFunc(controller, index, c)
}

func (controller Controller) LogAll(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session)
	go service.LogAll(collection)

}

func (controller Controller) LogByUserName(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session)
	go service.LogByUserName(collection, c)
}

func (controller Controller) ClearDatabase(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session)
	go service.ClearDatabase(collection)

}

func helperFunc(
	controller Controller,
	index int,
	c *gin.Context) {

	eventmap := controller.eventMap
	collection := utils.GetEventCollection(controller.session)
	eventType := eventmap[index][EVENT_TYPE].(string)
	go service.Processing(eventType, c, collection, eventmap)
}
