package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	// "log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) ProfileTransactionData() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")
		var ReqUserData []string

		err := c.BindJSON(&ReqUserData)
		if err != nil {
			LogError.LogError("🚀 ~ file: userProfileCartData.go ~ line 25 ~ returnfunc ~ err : ", err)
			
			c.JSON(http.StatusBadRequest, gin.H{"error": "cart list invalid"})
			return
		}
		fmt.Println("🚀 ~ file: userProfileCartData.go ~ line 26 ~ returnfunc ~ ReqUserData : ", ReqUserData)
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		// var UserDataDB model.UserData

		// err = H.Mongo.Collection(os.Getenv("USERDATA_COL")).FindOne(ctx, bson.M{"UserID": ReqUserData.UserID}).Decode(&UserDataDB)
		if err != nil {
			LogError.LogError("profiledata mongo find one error", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		}

		// CartData :=UserDataDB.UserCart

		var results []model.Gamedata
		if len(ReqUserData) != 0 {

			for _, v := range ReqUserData {
				var result model.Gamedata
			err := H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).FindOne(ctx, bson.D{{"GameID", v}}).Decode(&result)
			if err != nil {
            LogError.LogError("🚀 ~ file: userProfileCartData.go ~ line 43 ~ returnfunc ~ err : ", err)
				if err == mongo.ErrNoDocuments {
					c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
					return
				}
				c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
				return
			}
			results = append(results, result)
		}
	}


		

		// filter := bson.D{{"GameID", bson.D{{"$et", QueryId}}}}
		// opts := options.FindOne().SetSort(bson.D{})

		// var res bson.M
		// err := H.Mongo.Collection("GameData").FindOne(ctx, filter, opts).Decode(&res)

		// fmt.Println("🚀 ~ file: GetGameData.go ~ line 51 ~ returnfunc ~ res : ", res)

		// if err != nil {
		// 	log.Fatalln("❌ findOne : ", err)
		// }
		// fmt.Println("🚀✨ FindOne successful & result: ", res)

        fmt.Println("🚀 ~ file: userProfileCartData.go ~ line 71 ~ returnfunc ~ results : ", results)
		c.JSON(http.StatusOK, results)

	}
}
