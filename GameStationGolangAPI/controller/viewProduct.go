package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewProduct()gin.HandlerFunc{
	return func(c * gin.Context){
		c.JSON( http.StatusOK,"ViewProcut ok !!! " )
	}
}