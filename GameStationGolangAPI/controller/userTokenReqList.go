package controller

import (
	"app/LogError"
	"app/model"
	"fmt"

	// "app/model"
	"context"
	// "fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections) ProfileTokenReqList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")
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
		// findOptions := options.Find().SetProjection(bson.D{{"ReqList", 1},{"_id",0},})

		result := []model.UserMoney{}
		cursor, err := H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).Find(ctx, filter)

		if err != nil {
			LogError.LogError("🚀 ~ mogodb FindOne(ctx, filter, findOptions).Decode(&result) : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// var results []bson.D
		if err = cursor.All(context.TODO(), &result); err != nil {
			panic(err)
		}
		var Umoney model.UserMoney
		for _, r := range result {
			Umoney = r
		}
		fmt.Println("🚀 ~ file: userTokenReq.go ~ line 100 ~ returnfunc ~ Umoney.ReqList", Umoney.ReqList)
		c.JSON(http.StatusOK, Umoney.ReqList)

	}
}
