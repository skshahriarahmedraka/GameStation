package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) ProfileTokenReqList() gin.HandlerFunc {
	return func(c *gin.Context) {
		type ReqStruct struct {
			UserID string `json:"UserID" bson:"UserID"`
		}
		var ReqData ReqStruct

		err := c.BindJSON(&ReqData)
		if err != nil {
			LogError.LogError("❌🔥 error in c.BindJSON(&ReqData) ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		if ReqData.UserID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request userid "})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		//bson.M{"UserID": ReqData.UserID}
		filter := bson.M{"UserID": ReqData.UserID}
		findOptions := options.FindOne().SetProjection(bson.D{{"ReqList", 1}, {"Coin", 0}, {"UserID", 0}})

		result := []model.UserMoneyReq{}
		err = H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).FindOne(ctx, filter, findOptions).Decode(&result)
		fmt.Println("result", result)
		if err != nil {
			LogError.LogError("🚀 ~ mogodb FindOne(ctx, filter, findOptions).Decode(&result) : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)

		// authToken, err := c.Cookie("Auth1")
		// if authToken == "" {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthenticated User, cookie Auth1 not found"})
		// 	c.Abort()
		// 	return
		// }
		// claims, str := token.ValidateMoneyJWT(ReqData.JWT)
		// if str != "" {
		// 	LogError.LogError("🚀 ~ file: auth.go ~ line 22 ~ returnfunc ~ str : ", errors.New("jwt error"))
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Claims error"})
		// 	c.Abort()
		// 	return

		// }
		// c.Set("Email", claims.)
		// c.Set("UserID", claims.UserID)
		// c.Set("FirstName", claims.Name)

		// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		// defer cancel()
		// H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).FindOneAndUpdate(ctx, bson.M{"UserID": ReqData.UserID}, bson.M{"$inc": bson.M{"Amount": ""}})
	}
}
