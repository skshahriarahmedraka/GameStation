package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Instantbuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,"Instantbuy ok !!! ")
	}
}