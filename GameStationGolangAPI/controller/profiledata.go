package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (H *DatabaseCollections) UserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(http.StatusOK, "Login successful !!! ")
		//c.Param("profileid")
		c.JSON(http.StatusOK, gin.H{"profile": c.Param("profileid")})

	}
}
