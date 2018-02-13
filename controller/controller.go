package controller

import (
	"../service"
	"github.com/gin-gonic/gin"
)

func Usercommand(c *gin.Context) {
	service.Processing("usercommand", c)
}

func Systemevent(c *gin.Context) {
	service.Processing("systemevent", c)
}

func Quoteserver(c *gin.Context) {
	service.Processing("quoteserver", c)
}

func Accounttransaction(c *gin.Context) {
	service.Processing("accounttransaction", c)
}

func Errorevent(c *gin.Context) {
	service.Processing("errorevent", c)
}
