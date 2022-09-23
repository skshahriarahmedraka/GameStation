package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections)ViewProduct()gin.HandlerFunc{
	return func(c * gin.Context){
		c.JSON( http.StatusOK,"ViewProcut ok !!! " )
	}
}