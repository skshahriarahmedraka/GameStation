package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Deleteaddress() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(http.StatusOK,"Deleteaddress !!! ")
	}
}