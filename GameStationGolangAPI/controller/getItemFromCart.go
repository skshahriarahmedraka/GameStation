package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(http.StatusOK,"GetItemFromCart ok !!!")
	}
}