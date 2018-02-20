package controllers

import (
	"../service"
	"../utils"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

type (
	Controller struct {
		session *mgo.Session
	}
)

func NewController(s *mgo.Session) *Controller {
	return &Controller{s}
}

func (controller Controller) Usercommand(c *gin.Context) {

	newSession := controller.session.Clone()
	collection := utils.GetEventCollection(newSession)
	service.Processing("usercommand", c, collection)

}

func (controller Controller) Systemevent(c *gin.Context) {

	newSession := controller.session.Clone()
	collection := utils.GetEventCollection(newSession)
	service.Processing("systemevent", c, collection)

}

func (controller Controller) Quoteserver(c *gin.Context) {

	newSession := controller.session.Clone()
	collection := utils.GetEventCollection(newSession)
	service.Processing("quoteserver", c, collection)

}

func (controller Controller) Accounttransaction(c *gin.Context) {

	newSession := controller.session.Clone()
	collection := utils.GetEventCollection(newSession)
	service.Processing("accounttransaction", c, collection)

}

func (controller Controller) Errorevent(c *gin.Context) {

	newSession := controller.session.Clone()
	collection := utils.GetEventCollection(newSession)
	service.Processing("errorevent", c, collection)

}

func (controller Controller) LogAll(c *gin.Context) {

	newSession := controller.session.Copy()
	collection := utils.GetEventCollection(newSession)
	service.LogAll(collection)
	defer newSession.Close()

}

func (controller Controller) LogByUserName(c *gin.Context) {

	newSession := controller.session.Copy()
	collection := utils.GetEventCollection(newSession)
	service.LogByUserName(collection, c)
	defer newSession.Close()

}
