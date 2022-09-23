package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (H *DatabaseCollections)UserSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,"UserSearch ok...")
	}
}