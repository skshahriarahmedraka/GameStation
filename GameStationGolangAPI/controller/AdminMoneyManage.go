package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) AdminMoneyManagement() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		ctx,cancel := context.WithTimeout(context.Background(),time.Second*20)
		defer cancel()
		filter := bson.D{{}}
		cursor ,err := H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).Find(ctx, filter)
		if err!=nil {
        LogError.LogError("🚀 ~ file: AdminMoneyManage.go ~ line 28 ~ returnfunc ~ err : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var results []model.UserMoney
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		// var Umoney model.UserMoney
		// for _, r := range result {
		// 	Umoney = r
		// }

		c.JSON(http.StatusOK, results)

	}
}
