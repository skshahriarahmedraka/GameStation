package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Addaddress() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(http.StatusOK,"Addaddress ok !!!")
	}
}