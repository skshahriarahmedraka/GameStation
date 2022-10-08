package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections)AddtoCart() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"ok ":"Add to cart"})
	}
}