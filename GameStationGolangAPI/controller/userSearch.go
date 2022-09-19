package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func UserSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,"UserSearch ok...")
	}
}