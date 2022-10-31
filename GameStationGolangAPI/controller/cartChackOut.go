package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections) Cartcheckout() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.BindJSON()
		c.JSON(http.StatusOK, "Cartcheckout ok !!!!")

	}
}
