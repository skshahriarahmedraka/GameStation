package middleware

import (
	"app/LogError"
	"app/token"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {
		authToken, err := c.Cookie("Auth1")
		if authToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthenticated User, cookie Auth1 not found"})
			c.Abort()
			return
		}
		claims, str := token.ValidateJWT(authToken)
		if str != "" {
			LogError.LogError("🚀 ~ file: auth.go ~ line 22 ~ returnfunc ~ str : ", errors.New("jwt error"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Claims error"})
			c.Abort()
			return

		}
		c.Set("Email", claims.Email)
		c.Set("UserID", claims.UserID)
		c.Set("FirstName", claims.Name)
		//c.Set("LastName", claims.LastName)
		LogError.LogError("🚀 ~ file: auth.go ~ line 9 ~ returnfunc ~ err : ", err)
	}

}
