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

	collection := utils.GetEventCollection(controller.session.Clone())
	service.Processing("usercommand", c, collection)
}

func (controller Controller) Systemevent(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session.Clone())
	service.Processing("systemevent", c, collection)
}

func (controller Controller) Quoteserver(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session.Clone())
	service.Processing("quoteserver", c, collection)
}

func (controller Controller) Accounttransaction(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session.Clone())
	service.Processing("accounttransaction", c, collection)
}

func (controller Controller) Errorevent(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session.Clone())
	service.Processing("errorevent", c, collection)
}

func (controller Controller) LogAll(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session.Clone())
	service.LogAll(collection)

}

func (controller Controller) LogByUserName(c *gin.Context) {

	collection := utils.GetEventCollection(controller.session.Clone())
	service.LogByUserName(collection, c)
}
