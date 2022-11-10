package controller

import (
	"app/LogError"
	"app/token"
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections) ProfileMoneyReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		
		type ReqStruct struct {
			UserID string `json:"UserID" bson:"UserID"`
			JWT    string `json:"JWT" bson:"JWT"`
		}
		var ReqData ReqStruct

		err := c.BindJSON(&ReqData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		if ReqData.UserID == "" || ReqData.JWT == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		authToken, err := c.Cookie("Auth1")
		if authToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthenticated User, cookie Auth1 not found"})
			c.Abort()
			return
		}
		claims, str := token.ValidateMoneyJWT(authToken)
		if str != "" {
			LogError.LogError("🚀 ~ file: auth.go ~ line 22 ~ returnfunc ~ str : ", errors.New("jwt error"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Claims error"})
			c.Abort()
			return

		}
		c.Set("Email", claims.Email)
		c.Set("UserID", claims.UserID)
		c.Set("FirstName", claims.Name)

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).FindOneAndUpdate(ctx, bson.M{"UserID": ReqData.UserID}, bson.M{"$inc": bson.M{"Amount": ""}})
	}
}
