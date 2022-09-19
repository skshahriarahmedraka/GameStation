package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Edithomeaddress()gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,"Edithomeaddress ok !!!! ")
	}
}