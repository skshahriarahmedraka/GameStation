package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections)Deleteaddress() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(http.StatusOK,"Deleteaddress !!! ")
	}
}